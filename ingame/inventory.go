package ingame

import (
	"encoding/json"
	"ethereal-dreams/api"
	"fmt"
	"io"
	"net/http"

	"go.minekube.com/gate/pkg/edition/java/proxy"
)

const Overworld = "/f031752e-def7-4e28-82d2-7795d1a2e11a/"

func GrabInventory(MCplayer proxy.Player) ([]Player, error) {
	var (
		spigotAPI          api.APICallBuilder
		RequestedInventory []Player
	)

	spigotAPI.SetURL("http://209.222.97.128:27091/v1/players/" + MCplayer.ID().String() + Overworld + "inventory")
	spigotAPI.SetMethod("GET")
	spigotAPI.SetKey("MH1gP+qb1776")

	res, err := spigotAPI.Build().Call()
	if err != nil {
		fmt.Println("Failed to perform HTTP Client request. Error: \n", err)
	}

	defer res.Body.Close()

	_, err = api.ParseResults(res, &RequestedInventory)
	if err != nil {
		fmt.Println("Failed to parse HTTP Client request. Error: \n", err)
	}

	return RequestedInventory, err
}

func BurnCurrency(MCplayer proxy.Player, amnt int) error {
	var (
		spigotAPI api.APICallBuilder
	)

	spigotAPI.SetURL("http://209.222.97.128:27091/v1/server/exec")
	spigotAPI.SetPayload("command=minecraft:clear%20 " + MCplayer.Username() + " nether_star " + fmt.Sprint(amnt))
	spigotAPI.SetMethod("POST")
	spigotAPI.SetContentType("application/x-www-form-urlencoded")
	_, err := spigotAPI.Build().Call()

	return err
}

func GiveCurrency(MCplayer proxy.Player, amnt int) error {
	var (
		spigotAPI api.APICallBuilder
	)

	spigotAPI.SetURL("http://209.222.97.128:27091/v1/server/exec")
	spigotAPI.SetPayload("command=minecraft:give%20 " + MCplayer.Username() + " nether_star" + `{Tags:[ethereal,currency],display:{Name:'[{"text":"Dreams","italic":false,"color":"red"}]',Lore:['[{"text":"Ethereal Currency","italic":false,"color":"gold"}]']},Enchantments:[{}]} ` + fmt.Sprint(amnt))
	spigotAPI.SetMethod("POST")
	spigotAPI.SetContentType("application/x-www-form-urlencoded")
	_, err := spigotAPI.Build().Call()

	return err
}

func ValidateExchange(playerInv []Player, swapAmnt int) bool {
	for _, items := range playerInv {
		if items.ID == "minecraft:nether_star" {
			if items.Count >= swapAmnt {
				return true
			}
		}
	}

	return false
}

func ParseResults(res *http.Response, Results []PlayerInventory) ([]byte, error) {
	response, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Unable to parse API results. Error:\n", err)
	}
	err = json.Unmarshal(response, &Results)

	return response, err
}

type Player struct {
	PlayerInventory
}

type PlayerInventory struct {
	Slot  int    `json:"slot"`
	Count int    `json:"count"`
	ID    string `json:"id"`
}