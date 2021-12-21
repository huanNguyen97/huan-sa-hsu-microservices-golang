package B_BUS_layer

import (
	// Go importing owner
	"github.com/huan-hsu/microservices/layer/C_DAO_layer"
	"github.com/huan-hsu/microservices/DTO"
)

func ReadGamesBUS() []DTO.Game {
	games := C_DAO_layer.ReadGamesDAO()
	return games
}

func CreateGameBUS(game_posted *DTO.Game) bool {
	game_result := C_DAO_layer.CreateGameDAO(game_posted)
	return game_result
}

func ReadDetailBUS(id string) DTO.Game {
	game := C_DAO_layer.ReadGameDetailDAO(id)
	return game
}

func UpdateGameBUS(game_posted *DTO.Game) bool {
	game_result := C_DAO_layer.UpdateGameDAO(game_posted)
	return game_result
}

func DeleteGameBUS(id string) bool {
	game_result := C_DAO_layer.DeleteGameDAO(id)
	return game_result
}

func SearchGameBUS(key_name string) []DTO.Game {
	games := C_DAO_layer.SearchGameDAO(key_name)
	return games
}