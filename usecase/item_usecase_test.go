package usecase

func TestGetDoneItems(t *testing.T) {
	mockItemRepo := new(mocks.Item)
	var mockItem models.Item
	err := faker.FakeData(&mockItem)
	assert.NoError(t, err)

	mockListItem := make([]*models.Item, 0)
	mockListItem = append(mockListItem, &mockItem)
	mockItemRepo.On("Fetch", mock.AnythingOfType("string"), mock.AnythingOfType("int64")).Return(mockListItem, nil)
	u := ucase.NewItemUsecase(mockItemRepo)
	num := int64(1)
	description := "12"
	list, nextCursor, err := u.Fetch(cursor, num)
	cursorExpected := strconv.Itoa(int(mockItem.ID))
	assert.Equal(t, cursorExpected, nextCursor)
	assert.NotEmpty(t, nextCursor)
	assert.NoError(t, err)
	assert.Len(t, list, len(mockListItem))

	mockItemRepo.AssertCalled(t, "Fetch", mock.AnythingOfType("string"), mock.AnythingOfType("int64"))
}