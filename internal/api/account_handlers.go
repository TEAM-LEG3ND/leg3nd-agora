package api

import (
	"github.com/gofiber/fiber/v2"
	request "leg3nd-agora/internal/api/dto/request"
	response "leg3nd-agora/internal/api/dto/response"
	"leg3nd-agora/internal/core/ports"
)

type AccountHandlers struct {
	service ports.AccountService
}

func ProvideAccountHandlers(service ports.AccountService) *AccountHandlers {
	return &AccountHandlers{service: service}
}

func (co *AccountHandlers) CreateAccount(c *fiber.Ctx) error {
	var newAccountRequest *request.NewAccountRequest

	if err := c.BodyParser(&newAccountRequest); err != nil {
		return c.SendStatus(400)
	}

	newAccount, err := co.service.CreateAccount(c.Context(),
		newAccountRequest.Email,
		newAccountRequest.FullName,
		newAccountRequest.OAuthProvider,
	)
	if err != nil {
		return c.SendStatus(500)
	}

	newAccountResponse := &response.NewAccountResponse{Id: newAccount.Id}

	return c.JSON(newAccountResponse)
}

func (co *AccountHandlers) UpdateAccount(c *fiber.Ctx) error {
	var updateAccountRequest *request.UpdateAccountRequest

	if err := c.BodyParser(&updateAccountRequest); err != nil {
		return c.SendStatus(400)
	}

	updatedAccount, err := co.service.UpdateAccount(
		c.Context(),
		updateAccountRequest.Id,
		updateAccountRequest.Email,
		updateAccountRequest.FullName,
		updateAccountRequest.Nickname,
		updateAccountRequest.OAuthProvider,
		updateAccountRequest.Status,
	)
	if err != nil {
		return c.SendStatus(500)
	}

	updatedAccountResponse := &response.UpdatedAccountResponse{Id: updatedAccount.Id}

	return c.JSON(updatedAccountResponse)
}

func (co *AccountHandlers) FindAccountById(c *fiber.Ctx) error {
	var accountByIdRequest request.AccountByIdRequest

	err := c.ParamsParser(&accountByIdRequest)
	if err != nil {
		return err
	}

	account, err := co.service.FindAccountById(c.Context(), accountByIdRequest.Id)
	if err != nil {
		return err
	}

	accountResponse := response.AccountResponse{}.OfDomain(account)

	return c.JSON(accountResponse)
}

func (co *AccountHandlers) FindAccountByEmail(c *fiber.Ctx) error {
	var accountByEmailRequest request.AccountByEmailRequest

	err := c.ParamsParser(&accountByEmailRequest)
	if err != nil {
		return err
	}

	account, err := co.service.FindAccountByEmail(c.Context(), accountByEmailRequest.Email)
	if err != nil {
		return err
	}

	accountResponse := response.AccountResponse{}.OfDomain(account)

	return c.JSON(accountResponse)
}
