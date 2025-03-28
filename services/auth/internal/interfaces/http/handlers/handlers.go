package auth

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"net/http"

// 	"github.com/klimenkokayot/avito-go/pkg/models"
// 	service "github.com/klimenkokayot/avito-go/services/auth/internal/service"
// 	"github.com/sirupsen/logrus"
// )

// type AuthHandler struct {
// 	authService *service.AuthService
// }

// func NewAuthHandler() (*AuthHandler, error) {
// 	logrus.Info("Инициализация AuthHandler`a.")
// 	authService, err := service.NewAuthService()
// 	if err != nil {
// 		logrus.Error("Ошибка при создании AuthHandler`a.")
// 		return nil, fmt.Errorf("%w: %s", ErrCreateAuthHandler, err.Error())
// 	}
// 	logrus.Debug("Успешно создан AuthHandler.")
// 	return &AuthHandler{
// 		authService,
// 	}, nil
// }

// func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
// 	logrus.Info("Обработка запроса на регистрацию.")
// 	body, err := io.ReadAll(r.Body)
// 	if err != nil {
// 		logrus.Errorf("Неудачная регистрация пользователя (500)")
// 		w.WriteHeader(http.StatusInternalServerError)
// 		io.Writer(w).Write([]byte(fmt.Errorf("%w: %s", ErrReadBody, err.Error()).Error()))
// 		return
// 	}
// 	defer r.Body.Close()

// 	user := &models.User{}
// 	err = json.Unmarshal(body, user)
// 	if err != nil {
// 		logrus.Errorf("Неудачная регистрация пользователя (422)")
// 		w.WriteHeader(http.StatusUnprocessableEntity)
// 		io.Writer(w).Write([]byte(fmt.Errorf("%w: %s", ErrUnprocessibleEntity, err.Error()).Error()))
// 		return
// 	}

// 	err = h.authService.Register(user.Login, user.Password)
// 	if err != nil {
// 		logrus.Errorf("Пользователь существует (401)")
// 		w.WriteHeader(http.StatusUnauthorized)
// 		io.Writer(w).Write([]byte(fmt.Errorf("%w: %s", ErrRegisterProblem, err.Error()).Error()))
// 		return
// 	}

// 	logrus.Debugf("Пользователь успешно зарегистрирован: %s", user.Login)
// 	w.WriteHeader(http.StatusOK)
// }

// func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
// 	logrus.Info("Обработка запроса на вход.")
// 	body, err := io.ReadAll(r.Body)
// 	if err != nil {
// 		logrus.Errorf("Неудачный вход пользователя (500)")
// 		w.WriteHeader(http.StatusInternalServerError)
// 		io.Writer(w).Write([]byte(fmt.Errorf("%w: %s", ErrReadBody, err.Error()).Error()))
// 		return
// 	}
// 	defer r.Body.Close()

// 	user := &models.User{}
// 	err = json.Unmarshal(body, user)
// 	if err != nil {
// 		logrus.Errorf("Неудачная вход пользователя (422)")
// 		w.WriteHeader(http.StatusUnprocessableEntity)
// 		io.Writer(w).Write([]byte(fmt.Errorf("%w: %s", ErrUnprocessibleEntity, err.Error()).Error()))
// 		return
// 	}

// 	err = h.authService.Login(user.Login, user.Password)
// 	if err != nil {
// 		logrus.Errorf("Неправильный логин или пароль (401)")
// 		w.WriteHeader(http.StatusUnauthorized)
// 		io.Writer(w).Write([]byte(fmt.Errorf("%w: %s", ErrRegisterProblem, err.Error()).Error()))
// 		return
// 	}

// 	logrus.Debugf("Пользователь успешно вошел: %s", user.Login)
// 	w.WriteHeader(http.StatusOK)
// }
