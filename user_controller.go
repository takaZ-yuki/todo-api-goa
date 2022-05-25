package todoapi

import (
	"context"
	"log"
	usercontroller "todo-api/gen/user_controller"
	"todo-api/src/database"
	"todo-api/src/domain/model"

	"github.com/jinzhu/copier"
)

// user_controller service example implementation.
// The example methods log the requests and return zero values.
type userControllersrvc struct {
	logger *log.Logger
}

// NewUserController returns the user_controller service implementation.
func NewUserController(logger *log.Logger) usercontroller.Service {
	return &userControllersrvc{logger}
}

// ユーザ一覧の検索
func (s *userControllersrvc) GetUsers(ctx context.Context) (res []*usercontroller.User, err error) {
	s.logger.Print("userController.GetUsers")

	users := []model.User{}
	// ユーザー一覧を検索
	database.DB.Find(&users)

	// 返還用オブジェクト作成
	result_users := []*usercontroller.User{}
	// 検索結果を返還用オブジェクトへマッピング
	copier.Copy(&result_users, &users)

	return result_users, nil
}

// ユーザ検索
func (s *userControllersrvc) GetUser(ctx context.Context, p *usercontroller.GetUserPayload) (res *usercontroller.User, err error) {
	s.logger.Print("userController.GetUser")

	user := model.User{ID: p.ID}
	// ユーザー検索
	database.DB.Take(&user)

	// 返還用オブジェクト作成
	res = &usercontroller.User{}
	// 検索結果を返還用オブジェクトへマッピング
	copier.Copy(&res, &user)

	return res, nil
}

// ユーザ更新
func (s *userControllersrvc) UpdateUser(ctx context.Context, p *usercontroller.User) (err error) {
	s.logger.Print("userController.UpdateUser")
	user := model.User{ID: *p.ID}
	// ユーザー検索
	database.DB.Take(&user)
	copier.Copy(&user, &p)

	// ユーザー更新処理
	database.DB.Save(&user)
	return
}

// ユーザ登録
func (s *userControllersrvc) CreateUser(ctx context.Context, p *usercontroller.User) (err error) {
	s.logger.Print("userController.CreateUser")
	user := model.User{}
	copier.Copy(&user, &p)

	// ユーザー登録処理
	database.DB.Save(&user)
	return
}

// ユーザ削除
func (s *userControllersrvc) DeleteUser(ctx context.Context, p *usercontroller.DeleteUserPayload) (err error) {
	s.logger.Print("userController.DeleteUser")
	// ユーザー削除処理
	database.DB.Delete(&model.User{}, p.ID)
	return
}
