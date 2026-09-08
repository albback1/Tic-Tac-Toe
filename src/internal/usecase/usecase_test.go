package usecase

import (
	"testing"
	"tictac/internal/adapter/inmemory"
	"tictac/internal/domain"

	"github.com/google/uuid"
)

func Test_GameOverCheck_winX(t *testing.T) {

	var repo inmemory.InMemory

	var service *GameServiceItml
	service = NewGameServiceItml(&repo)

	uuid := uuid.New()

	var testField [12]*domain.Game
	for i := 0; i < 12; i++ {
		testField[i] = domain.NewGame(3, uuid)
	}

	testField[0].Field = [][]domain.Cell{
		{1, 1, 1},
		{0, 0, 0},
		{0, 0, 0}}
	testField[1].Field = [][]domain.Cell{
		{0, 0, 0},
		{1, 1, 1},
		{0, 0, 0}}
	testField[2].Field = [][]domain.Cell{
		{0, 0, 0},
		{1, 1, 1},
		{0, 0, 0}}
	testField[3].Field = [][]domain.Cell{
		{0, 0, 0},
		{0, 0, 0},
		{1, 1, 1}}
	testField[4].Field = [][]domain.Cell{
		{1, 0, 0},
		{1, 0, 0},
		{1, 0, 0}}
	testField[5].Field = [][]domain.Cell{
		{0, 1, 0},
		{0, 1, 0},
		{0, 1, 0}}
	testField[6].Field = [][]domain.Cell{
		{0, 0, 1},
		{0, 0, 1},
		{0, 0, 1}}
	testField[7].Field = [][]domain.Cell{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1}}
	testField[8].Field = [][]domain.Cell{
		{0, 0, 1},
		{0, 1, 0},
		{1, 0, 0}}
	testField[9].Field = [][]domain.Cell{
		{1, 1, 1},
		{0, 2, 2},
		{2, 1, 0}}
	testField[10].Field = [][]domain.Cell{
		{0, 0, 1},
		{0, 1, 2},
		{1, 0, 2}}
	testField[11].Field = [][]domain.Cell{
		{1, 1, 1},
		{2, 2, 0},
		{0, 1, 0}}

	for i, x := range testField {
		result := service.GameOverCheck(x)
		t.Logf("tectcase %d", i)
		if result != domain.XWin {
			t.Errorf("Incorrect result. Expect: %d, res: %d", 1, result)
		}

	}
}

func Test_GameOverCheck_winO(t *testing.T) {

	var repo inmemory.InMemory

	var service *GameServiceItml
	service = NewGameServiceItml(&repo)

	uuid := uuid.New()

	var testField [12]*domain.Game
	for i := 0; i < 12; i++ {
		testField[i] = domain.NewGame(3, uuid)
	}

	testField[0].Field = [][]domain.Cell{
		{2, 2, 2},
		{0, 0, 0},
		{0, 0, 0}}
	testField[1].Field = [][]domain.Cell{
		{0, 0, 0},
		{2, 2, 2},
		{0, 0, 0}}
	testField[2].Field = [][]domain.Cell{
		{2, 0, 1},
		{2, 2, 1},
		{0, 0, 2}}
	testField[3].Field = [][]domain.Cell{
		{0, 0, 0},
		{0, 0, 0},
		{2, 2, 2}}
	testField[4].Field = [][]domain.Cell{
		{2, 0, 0},
		{2, 0, 0},
		{2, 0, 0}}
	testField[5].Field = [][]domain.Cell{
		{0, 2, 0},
		{0, 2, 0},
		{0, 2, 0}}
	testField[6].Field = [][]domain.Cell{
		{0, 0, 2},
		{0, 0, 2},
		{0, 0, 2}}
	testField[7].Field = [][]domain.Cell{
		{2, 0, 0},
		{0, 2, 0},
		{0, 0, 2}}
	testField[8].Field = [][]domain.Cell{
		{0, 0, 2},
		{0, 2, 0},
		{2, 0, 0}}
	testField[9].Field = [][]domain.Cell{
		{2, 2, 2},
		{0, 1, 1},
		{1, 2, 0}}
	testField[10].Field = [][]domain.Cell{
		{0, 0, 2},
		{0, 2, 1},
		{2, 0, 1}}
	testField[11].Field = [][]domain.Cell{
		{2, 2, 2},
		{1, 1, 0},
		{0, 2, 0}}

	for i, x := range testField {
		result := service.GameOverCheck(x)
		t.Logf("tectcase %d", i)
		if result != domain.OWin {
			t.Errorf("Incorrect result. Expect: %d, res: %d", 2, result)
		}
	}
}

func Test_GameOverCheck_Ongoing(t *testing.T) {

	var repo inmemory.InMemory

	var service *GameServiceItml
	service = NewGameServiceItml(&repo)

	uuid := uuid.New()

	var testField [12]*domain.Game
	for i := 0; i < 12; i++ {
		testField[i] = domain.NewGame(3, uuid)
	}

	testField[0].Field = [][]domain.Cell{
		{0, 2, 2},
		{0, 0, 0},
		{0, 0, 0}}
	testField[1].Field = [][]domain.Cell{
		{0, 0, 0},
		{2, 0, 2},
		{0, 0, 0}}
	testField[2].Field = [][]domain.Cell{
		{2, 0, 1},
		{2, 0, 1},
		{0, 0, 2}}
	testField[3].Field = [][]domain.Cell{
		{0, 0, 0},
		{0, 0, 0},
		{2, 0, 2}}
	testField[4].Field = [][]domain.Cell{
		{2, 0, 0},
		{0, 0, 0},
		{2, 0, 0}}
	testField[5].Field = [][]domain.Cell{
		{0, 2, 0},
		{0, 0, 0},
		{0, 2, 0}}
	testField[6].Field = [][]domain.Cell{
		{0, 0, 2},
		{0, 0, 0},
		{0, 0, 2}}
	testField[7].Field = [][]domain.Cell{
		{2, 0, 0},
		{0, 0, 0},
		{0, 0, 2}}
	testField[8].Field = [][]domain.Cell{
		{0, 0, 2},
		{0, 0, 0},
		{2, 0, 0}}
	testField[9].Field = [][]domain.Cell{
		{2, 0, 2},
		{2, 1, 1},
		{1, 2, 0}}
	testField[10].Field = [][]domain.Cell{
		{0, 0, 2},
		{0, 0, 1},
		{2, 0, 1}}
	testField[11].Field = [][]domain.Cell{
		{2, 0, 2},
		{1, 1, 0},
		{0, 2, 0}}

	for i, x := range testField {
		result := service.GameOverCheck(x)
		if result != domain.Ongoing {
			t.Errorf("Incorrect result. Expect: %d, res: %d, testcase %d", 0, result, i)
		}
	}
}

func Test_GameOverCheck_Draw(t *testing.T) {

	var repo inmemory.InMemory

	var service *GameServiceItml
	service = NewGameServiceItml(&repo)

	uuid := uuid.New()

	var testField [12]*domain.Game
	for i := 0; i < 12; i++ {
		testField[i] = domain.NewGame(3, uuid)
	}

	testField[0].Field = [][]domain.Cell{
		{1, 1, 2},
		{2, 2, 1},
		{1, 2, 1}}
	testField[1].Field = [][]domain.Cell{
		{2, 1, 2},
		{2, 1, 2},
		{1, 2, 1}}
	testField[2].Field = [][]domain.Cell{
		{2, 1, 2},
		{2, 1, 1},
		{1, 2, 2}}
	testField[3].Field = [][]domain.Cell{
		{2, 1, 2},
		{2, 1, 1},
		{1, 2, 2}}
	testField[4].Field = [][]domain.Cell{
		{2, 2, 1},
		{1, 2, 2},
		{2, 1, 1}}
	testField[5].Field = [][]domain.Cell{
		{1, 2, 2},
		{2, 1, 1},
		{1, 2, 2}}
	testField[6].Field = [][]domain.Cell{
		{2, 1, 2},
		{2, 1, 1},
		{1, 2, 2}}
	testField[7].Field = [][]domain.Cell{
		{2, 1, 2},
		{2, 1, 2},
		{1, 2, 1}}
	testField[8].Field = [][]domain.Cell{
		{1, 1, 2},
		{2, 1, 1},
		{1, 2, 2}}
	testField[9].Field = [][]domain.Cell{
		{2, 1, 2},
		{2, 1, 1},
		{1, 2, 2}}
	testField[10].Field = [][]domain.Cell{
		{1, 1, 2},
		{2, 2, 1},
		{1, 1, 2}}
	testField[11].Field = [][]domain.Cell{
		{2, 1, 2},
		{1, 1, 2},
		{1, 2, 1}}

	for i, x := range testField {
		result := service.GameOverCheck(x)
		if result != domain.Draw {
			t.Errorf("Incorrect result. Expect: %d, res: %d, testcase %d", 3, result, i)
		}
	}
}

func Test_validateFieldx_tru(t *testing.T) {
	var repo inmemory.InMemory

	var service *GameServiceItml
	service = NewGameServiceItml(&repo)

	var domainField [12]*domain.Game
	for i := 0; i < 12; i++ {
		uuid := uuid.New()
		domainField[i] = domain.NewGame(3, uuid)
	}

	domainField[0].Field = [][]domain.Cell{
		{1, 1, 2},
		{2, 2, 1},
		{1, 2, 1}}
	domainField[1].Field = [][]domain.Cell{
		{2, 1, 2},
		{2, 1, 2},
		{1, 2, 1}}
	domainField[2].Field = [][]domain.Cell{
		{2, 1, 2},
		{2, 1, 1},
		{1, 2, 2}}
	domainField[3].Field = [][]domain.Cell{
		{2, 1, 2},
		{2, 1, 1},
		{1, 2, 2}}
	domainField[4].Field = [][]domain.Cell{
		{2, 2, 1},
		{1, 2, 2},
		{2, 1, 1}}
	domainField[5].Field = [][]domain.Cell{
		{1, 2, 2},
		{2, 1, 1},
		{1, 2, 2}}
	domainField[6].Field = [][]domain.Cell{
		{2, 1, 2},
		{2, 1, 1},
		{1, 2, 2}}
	domainField[7].Field = [][]domain.Cell{
		{2, 1, 2},
		{2, 1, 2},
		{1, 2, 1}}
	domainField[8].Field = [][]domain.Cell{
		{1, 1, 2},
		{2, 1, 1},
		{1, 2, 2}}
	domainField[9].Field = [][]domain.Cell{
		{2, 1, 2},
		{2, 1, 1},
		{1, 2, 2}}
	domainField[10].Field = [][]domain.Cell{
		{1, 1, 2},
		{2, 2, 1},
		{1, 1, 2}}
	domainField[11].Field = [][]domain.Cell{
		{2, 1, 2},
		{1, 1, 2},
		{1, 2, 1}}

	for i := 0; i < len(domainField); i++ {
		service.Repo.SaveGame(domainField[i])
	}

	for i, x := range domainField {
		result, _ := service.ValidateField(x)
		if !result {
			t.Errorf("Incorrect result. testcase %d", i)
		}
	}

}

func Test_validateFieldx_false(t *testing.T) {
	var repo inmemory.InMemory

	var service *GameServiceItml
	service = NewGameServiceItml(&repo)

	var domainField [12]*domain.Game
	var newField [12]*domain.Game
	for i := 0; i < 12; i++ {
		uuid := uuid.New()
		domainField[i] = domain.NewGame(3, uuid)
		newField[i] = domain.NewGame(3, uuid)
	}

	domainField[0].Field = [][]domain.Cell{
		{1, 1, 2},
		{2, 2, 1},
		{1, 2, 1}}
	domainField[1].Field = [][]domain.Cell{
		{2, 1, 2},
		{2, 1, 2},
		{1, 2, 1}}
	domainField[2].Field = [][]domain.Cell{
		{2, 1, 2},
		{2, 1, 1},
		{1, 2, 2}}
	domainField[3].Field = [][]domain.Cell{
		{2, 1, 2},
		{2, 1, 1},
		{1, 2, 2}}
	domainField[4].Field = [][]domain.Cell{
		{2, 2, 1},
		{1, 2, 2},
		{2, 1, 1}}
	domainField[5].Field = [][]domain.Cell{
		{1, 2, 2},
		{2, 1, 1},
		{1, 2, 2}}
	domainField[6].Field = [][]domain.Cell{
		{2, 1, 2},
		{2, 1, 1},
		{1, 2, 2}}
	domainField[7].Field = [][]domain.Cell{
		{2, 1, 2},
		{2, 1, 2},
		{1, 2, 1}}
	domainField[8].Field = [][]domain.Cell{
		{1, 1, 2},
		{2, 1, 1},
		{1, 2, 2}}
	domainField[9].Field = [][]domain.Cell{
		{2, 1, 2},
		{2, 1, 1},
		{1, 2, 2}}
	domainField[10].Field = [][]domain.Cell{
		{1, 1, 2},
		{2, 2, 1},
		{1, 1, 2}}
	domainField[11].Field = [][]domain.Cell{
		{2, 1, 2},
		{1, 1, 2},
		{1, 2, 1}}

	for i := 0; i < len(domainField); i++ {
		service.Repo.SaveGame(domainField[i])
	}

	newField[0].Field = [][]domain.Cell{
		{1, 2, 2},
		{2, 2, 1},
		{1, 2, 1}}
	newField[1].Field = [][]domain.Cell{
		{2, 2, 2},
		{2, 1, 2},
		{1, 2, 1}}
	newField[2].Field = [][]domain.Cell{
		{2, 2, 2},
		{2, 1, 1},
		{1, 2, 2}}
	newField[3].Field = [][]domain.Cell{
		{2, 2, 2},
		{2, 1, 1},
		{1, 2, 2}}
	newField[4].Field = [][]domain.Cell{
		{2, 1, 1},
		{1, 2, 2},
		{2, 1, 1}}
	newField[5].Field = [][]domain.Cell{
		{1, 1, 2},
		{2, 1, 1},
		{1, 2, 2}}
	newField[6].Field = [][]domain.Cell{
		{2, 2, 2},
		{2, 1, 1},
		{1, 2, 2}}
	newField[7].Field = [][]domain.Cell{
		{2, 2, 2},
		{2, 1, 2},
		{1, 2, 1}}
	newField[8].Field = [][]domain.Cell{
		{1, 2, 2},
		{2, 1, 1},
		{1, 2, 2}}
	newField[9].Field = [][]domain.Cell{
		{2, 2, 2},
		{2, 1, 1},
		{1, 2, 2}}
	newField[10].Field = [][]domain.Cell{
		{1, 2, 2},
		{2, 2, 1},
		{1, 1, 2}}
	newField[11].Field = [][]domain.Cell{
		{2, 2, 2},
		{1, 1, 2},
		{1, 2, 1}}

	for i, x := range newField {
		result, _ := service.ValidateField(x)
		if result {
			t.Errorf("Incorrect result. testcase %d", i)
		}
	}

}

func Test_NextMoveMinimax(t *testing.T) {
	var repo inmemory.InMemory

	var service *GameServiceItml
	service = NewGameServiceItml(&repo)

	var domainField [5]*domain.Game

	for i := 0; i < 5; i++ {
		uuid := uuid.New()
		domainField[i] = domain.NewGame(3, uuid)
	}

	domainField[0].Field = [][]domain.Cell{
		{1, 1, 0},
		{2, 2, 0},
		{0, 0, 0}}

	var x, y = service.NextMoveMinimax(domainField[0], domain.X)

	if x != 2 || y != 0 {
		t.Errorf("Incorrect result. testcase %d", 0)
	}

	domainField[1].Field = [][]domain.Cell{
		{2, 2, 0},
		{1, 0, 0},
		{0, 0, 0}}
	x, y = service.NextMoveMinimax(domainField[1], domain.X)

	if x != 2 || y != 0 {
		t.Errorf("Incorrect result. testcase %d", 1)
	}

	domainField[2].Field = [][]domain.Cell{
		{1, 0, 0},
		{2, 2, 0},
		{0, 0, 0}}
	x, y = service.NextMoveMinimax(domainField[2], domain.X)

	if x != 2 || y != 1 {
		t.Errorf("Incorrect result. testcase %d", 2)
		t.Errorf("%d %d", x, y)
	}

	domainField[3].Field = [][]domain.Cell{
		{2, 0, 0},
		{1, 2, 0},
		{0, 0, 1}}
	x, y = service.NextMoveMinimax(domainField[3], domain.X)

	if x == 2 || y == 1 {
		t.Errorf("Incorrect result. testcase %d", 3)
		t.Errorf("%d %d", x, y)
	}

	domainField[4].Field = [][]domain.Cell{
		{1, 0, 0},
		{2, 1, 0},
		{2, 0, 0}}
	x, y = service.NextMoveMinimax(domainField[4], domain.X)

	if x != 2 || y != 2 {
		t.Errorf("Incorrect result. testcase %d", 4)
		t.Errorf("%d %d", x, y)
	}
}
