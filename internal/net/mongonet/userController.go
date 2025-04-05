package mongonet

//func GetAllUsersHandler(c echo.Context) error {
//	client := c.Get("mongoClient").(*mongonet.Client)
//	usersls, err := users.GetAll(client)
//	if err != nil {
//		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
//	}
//	if usersls == nil {
//		usersls = []models.User{}
//	}
//	return c.JSON(http.StatusOK, usersls)
//}
//
//func PostUserHandler(c echo.Context) error {
//	var newUser models.UserWithoutId
//	client := c.Get("mongoClient").(*mongonet.Client)
//	if err := c.Bind(&newUser); err != nil {
//		return c.JSON(http.StatusBadRequest, map[string]string{"error": "INVALID JSON"})
//	}
//
//	err := users.Create(client, models.User{
//		Id:       uuid.NewString(),
//		Login:    newUser.Login,
//		Password: newUser.Password,
//	})
//	if err != nil {
//		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
//	}
//
//	return c.JSON(http.StatusOK, map[string]string{"message": "user post successfully"})
//}
//
//func PutUserHandler(c echo.Context) error {
//	client := c.Get("mongoClient").(*mongonet.Client)
//
//	var ud models.UserWithoutId
//	if err := c.Bind(&ud); err != nil {
//		return c.JSON(http.StatusBadRequest, map[string]string{"error": "INVALID JSON"})
//	}
//
//	id := c.QueryParam("id")
//	if id == "" {
//		return c.JSON(http.StatusBadRequest, map[string]string{"error": "id parameter is required"})
//	}
//
//
//
//	if err := users.Update(client, id, ud); err != nil {
//		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
//	}
//
//	return c.JSON(http.StatusOK, map[string]string{"message": "user update successfully"})
//}
//
//func DeleteUserHandler(c echo.Context) error {
//	client := c.Get("mongoClient").(*mongonet.Client)
//
//	id := c.QueryParam("id")
//	if id == "" {
//		return c.JSON(http.StatusBadRequest, map[string]string{"error": "id parameter is required"})
//	}
//
//	if err := users.Delete(client, id); err != nil {
//		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
//	}
//
//	return c.JSON(http.StatusOK, map[string]string{"message": "user delete successfully"})
//}
