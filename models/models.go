package models

import (
	"errors"
)

// Token balance response structure
type TokenBalance struct {
	TokenAddress string `json:"token_address"`
	Name         string `json:"name"`
	Symbol       string `json:"symbol"`
	Logo         string `json:"logo"`
	Thumbnail    string `json:"thumbnail"`
	Decimals     int    `json:"decimals"`
	Balance      string `json:"balance"`
	PossibleSpam bool   `json:"possible_spam"`
}

// Input file data structure for token balances
type TokenFile struct {
	Address string
	Chain   string
	Block   string
}

// Native balance response structure
type NativeBalance struct {
	Balance string `json:"balance"`
}

// Incoming request body struct
type Request struct {
	Address   string `json:"address"`
	Chain     string `json:"chain"`
	Date      string `json:"date"`
	Timestamp string `json:"timestamp"`
}

// Response sent to client
type Response struct {
	Address          string               `json:"account_address"`
	Chain            string               `json:"chain"`
	BlockNumber      int                  `json:"block_number"`
	Asset            string               `json:"native_asset"`
	AssetName        string               `json:"native_asset_name"`
	NativeCheckerUrl string               `json:"native_checker_url"`
	Balance          float64              `json:"native_balance"`
	ERC20Tokens      []ERC20TokenResponse `json:"erc20_tokens"`
}

// Response sent to client (token array)
type ERC20TokenResponse struct {
	ERC20TokenName            string `json:"erc20_token_name"`
	ERC20TokenSymbol          string `json:"erc20_token_symbol"`
	ERC20TokenContractAddress string `json:"erc20_token_contract_address"`
	ERC20TokenBalance         string `json:"erc20_token_balance"`
	ERC20TokenCheckerUrl      string `json:"erc20_token_checker_url"`
	ERC20PossibleSpam         bool   `json:"erc20_possible_spam"`
}

const (
	// ERC20 Token balance checker urls for auditability
	EthereumTokenChecker   = "https://etherscan.io/tokencheck-tool"
	ArbitrumTokenChecker   = "https://arbiscan.io/tokencheck-tool"
	PolygonTokenChecker    = "https://polygonscan.com/tokencheck-tool"
	BSCTokenChecker        = "https://bscscan.com/tokencheck-tool"
	FantomTokenChecker     = "https://ftmscan.com/tokencheck-tool"
	AvalancheTokenChecker  = "https://snowtrace.io/tokencheck-tool"
	CronosTokenChecker     = "https://cronoscan.com/tokencheck-tool"
	EthereumNativeChecker  = "https://etherscan.io/balancecheck-tool"
	ArbitrumNativeChecker  = "https://arbiscan.io/balancecheck-tool"
	PolygonNativeChecker   = "https://polygonscan.com/balancecheck-tool"
	BSCNativeChecker       = "https://bscscan.com/balancecheck-tool"
	FantomNativeChecker    = "https://ftmscan.com/balancecheck-tool"
	AvalancheNativeChecker = "https://snowtrace.io/balancecheck-tool"
	CronosNativeChecker    = "https://cronoscan.com/balancecheck-tool"

	// Native tokens
	ETH   = "ETH"
	ARB   = "ETH"
	BSC   = "BNB"
	FTM   = "FTM"
	MATIC = "MATIC"
	AVAX  = "AVAX"
	CRO   = "CRO"
)

var (
	// Chain variables for Moralis API
	Ethereum          = "eth"
	Arbitrum          = "arbitrum"
	Polygon           = "polygon"
	BinanceSmartChain = "bsc"
	Fantom            = "fantom"
	Avalanche         = "avalanche"
	Cronos            = "cronos"
)

func DetermineChain(chain string) (string, error) {
	var blockchain string

	if chain == "eth" || chain == "ethereum" || chain == "Ethereum" || chain == "ETH" || chain == "Eth" {
		blockchain = Ethereum
	} else if chain == "polygon" || chain == "matic" || chain == "Polygon" || chain == "MATIC" || chain == "Matic" {
		blockchain = Polygon
	} else if chain == "arbitrum" || chain == "Arbitrum" || chain == "arb" {
		blockchain = Arbitrum
	} else if chain == "bsc" || chain == "binance" || chain == "binance smart chain" || chain == "bnb chain" || chain == "bnb" || chain == "BNB" || chain == "Binance Smart Chain" || chain == "BSC" {
		blockchain = BinanceSmartChain
	} else if chain == "ftm" || chain == "fantom" || chain == "FTM" || chain == "Fantom" {
		blockchain = Fantom
	} else if chain == "cro" || chain == "CRO" || chain == "cronos" || chain == "Cronos" {
		blockchain = Cronos
	} else if chain == "avax" || chain == "avalanche" || chain == "AVAX" {
		blockchain = Avalanche
	} else {
		err := errors.New("did you make a typo? if not, then that blockchain is not supported. please use one of the supported chains: ethereum, arbitrum, polygon, binance smart chain, fantom, avalanche, cronos")
		return "", err
	}

	return blockchain, nil
}

// Determine info such as native token checker url, token name to be returned to user
func ReturnNativeInfo(chain string) (string, string, string, error) {
	// Store asset symbol
	var asset string

	// Store native token checker url
	var checkerUrl string

	// Store token name
	var tokenName string

	blockchain, err := DetermineChain(chain)
	if err != nil {
		return "", "", "", err
	}

	if blockchain == Ethereum {
		asset = ETH
		checkerUrl = EthereumNativeChecker
		tokenName = "Ethereum"
	} else if blockchain == Arbitrum {
		asset = ARB
		checkerUrl = ArbitrumNativeChecker
		tokenName = "Ethereum"
	} else if blockchain == BinanceSmartChain {
		asset = BSC
		checkerUrl = BSCNativeChecker
		tokenName = "Binance Coin"
	} else if blockchain == Fantom {
		asset = FTM
		checkerUrl = FantomNativeChecker
		tokenName = "Fantom"
	} else if blockchain == Avalanche {
		asset = AVAX
		checkerUrl = AvalancheNativeChecker
		tokenName = "Avalanche"
	} else if blockchain == Polygon {
		asset = MATIC
		checkerUrl = PolygonNativeChecker
		tokenName = "Polygon (MATIC)"
	} else if blockchain == Cronos {
		asset = CRO
		checkerUrl = CronosNativeChecker
		tokenName = "Cronos"
	} else {
		return "", "", "", errors.New("native information cannot be determined")
	}

	return asset, checkerUrl, tokenName, nil
}

func ReturnERC20TokenChecker(chain string) (string, error) {
	var checkerUrl string

	blockchain, err := DetermineChain(chain)
	if err != nil {
		return "", err
	}

	if blockchain == Ethereum {
		checkerUrl = EthereumTokenChecker
	} else if blockchain == Arbitrum {
		checkerUrl = ArbitrumTokenChecker
	} else if blockchain == BinanceSmartChain {
		checkerUrl = BSCTokenChecker
	} else if blockchain == Fantom {
		checkerUrl = FantomTokenChecker
	} else if blockchain == Avalanche {
		checkerUrl = AvalancheTokenChecker
	} else if blockchain == Polygon {
		checkerUrl = PolygonTokenChecker
	} else if blockchain == Cronos {
		checkerUrl = CronosTokenChecker
	} else {
		return "", errors.New("erc20 token checker url cannot be determined")
	}

	return checkerUrl, nil
}
