package main

import (
	"context"
	"math/big"
	"net/http"
	"os"

	"github.com/ElioenaiFerrari/blockchain/src/api"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	godotenv.Load()
}

func main() {
	client, err := ethclient.DialContext(context.Background(), os.Getenv("BLOCKCHAIN_URL"))

	if err != nil {
		panic(err)
	}

	defer client.Close()

	pvk, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))

	if err != nil {
		panic(err)
	}

	addr := crypto.PubkeyToAddress(pvk.PublicKey)

	nonce, err := client.PendingNonceAt(context.Background(), addr)

	if err != nil {
		panic(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())

	if err != nil {
		panic(err)
	}

	gasLimit := uint64(3000000)

	chainID, err := client.NetworkID(context.Background())

	if err != nil {
		panic(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(pvk, chainID)

	if err != nil {
		panic(err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	auth.GasLimit = gasLimit
	auth.GasPrice = gasPrice

	addr, _, _, err = api.DeployNFT(auth, client)

	if err != nil {
		panic(err)
	}

	conn, err := api.NewNFT(addr, client)

	if err != nil {
		panic(err)
	}

	e := echo.New()
	r := e.Group("/api/v1")

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())
	e.Use(middleware.CORS())

	r.POST("/tokens", func(c echo.Context) error {
		type Request struct {
			ImageURL string `json:"image_url"`
			Name     string `json:"name"`
			Price    uint64 `json:"price"`
		}

		var request Request

		if err := c.Bind(&request); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		if request.ImageURL == "" {
			return c.JSON(http.StatusBadRequest, "image_url is required")
		}

		if request.Name == "" {
			return c.JSON(http.StatusBadRequest, "name is required")
		}

		if request.Price == 0 {
			return c.JSON(http.StatusBadRequest, "price is required")
		}

		auth.Nonce = auth.Nonce.Add(auth.Nonce, big.NewInt(1))

		reply, err := conn.Mint(auth, uuid.New().String(), request.Name, request.Price, request.ImageURL)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, reply)
	})

	r.GET("/tokens", func(c echo.Context) error {
		tokens, err := conn.GetTokens(&bind.CallOpts{})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, tokens)

	})

	r.PUT("/tokens", func(c echo.Context) error {
		type Request struct {
			TokenID string `json:"token_id"`
			To      string `json:"to"`
		}

		var request Request

		if err := c.Bind(&request); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		if request.TokenID == "" {
			return c.JSON(http.StatusBadRequest, "id is required")
		}

		if request.To == "" {
			return c.JSON(http.StatusBadRequest, "to is required")
		}

		auth.Nonce = auth.Nonce.Add(auth.Nonce, big.NewInt(1))

		reply, err := conn.Send(auth, request.TokenID, common.HexToAddress(request.To))

		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, reply)
	})

	e.Logger.Fatal(e.Start(":3000"))
}
