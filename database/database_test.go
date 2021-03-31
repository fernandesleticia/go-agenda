import (
	"testing"
	"context"
	"time"

	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestCreateItem(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)		
	}
    
	defer db.Close()
	rows := sqlmock.NewRows([]string{"id", "Description", "Done"}).
	AddRow(1, "Buy an xbox", false)

	query := "SELECT id, Description, Done FROM items WHERE id=\\?"

	prep := mock.ExpectPrepare(query)
	uitemID := int64(1)
	prep.ExpectQuery().WithArgs(itemID).WillReturnRows(rows)

	a := database.MysqlInstance
    
	anItem, err := a.GetItemByID(num)  
	assert.NoError(t, err) 
	assert.NotNil(t, anArticle)
}