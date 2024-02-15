import { fireEvent, render, screen, waitFor } from "@testing-library/react";
import axios from "axios";
import MockAdapter from "axios-mock-adapter";
import SignIn from "./SignIn";

const mockHistoryPush = jest.fn();

const mockHistory = {
  push: mockHistoryPush,
};

jest.mock("react-router", () => {
  const originalModule = jest.requireActual("react-router");
  return {
    ...originalModule,
    useHistory: () => mockHistory,
  };
});

const mock = new MockAdapter(axios);
const data = { response: true };

describe("Sign In", () => {
  it("should have sign in page", () => {
    render(<SignIn />);
    const linkElement = screen.getByTestId("sign-in");
    expect(linkElement).toBeInTheDocument();
  });

  it("should render username input", () => {
    render(<SignIn />);
    const linkElement = screen.getByTestId("username-input");
    expect(linkElement).toBeInTheDocument();
  });

  it("should able to enter username input", () => {
    render(<SignIn />);
    const userNameInput = screen.getByTestId("username-input");
    fireEvent.change(userNameInput, { target: { value: "tushar" } });
    expect(userNameInput).toHaveValue("tushar");
  });

  it("should render password input", () => {
    render(<SignIn />);
    const passwordInput = screen.getByTestId("password-input");
    fireEvent.change(passwordInput, { target: { value: "tushar" } });
    expect(passwordInput).toHaveValue("tushar");
    expect(passwordInput).toBeInTheDocument();
  });

  it("should able to enter password input", () => {
    render(<SignIn />);
    const passwordInput = screen.getByTestId("password-input");
    fireEvent.change(passwordInput, { target: { value: "tushar" } });
    expect(passwordInput).toHaveValue("tushar");
  });

  it("should render submit btn input", () => {
    render(<SignIn />);
    const linkElement = screen.getByTestId("submit-btn");
    expect(linkElement).toBeInTheDocument();
  });

  it("should disable the submit button when username and password is not entered", () => {
    render(<SignIn />);
    const submitButton = screen.getByTestId("submit-btn");
    expect(submitButton).toBeDisabled();
  });

  it("should enabled the submit button when username and password is entered", () => {
    render(<SignIn />);
    const userNameInput = screen.getByTestId("username-input");
    fireEvent.change(userNameInput, { target: { value: "tushar" } });
    const passwordInput = screen.getByTestId("password-input");
    fireEvent.change(passwordInput, { target: { value: "tushar" } });
    const submitButton = screen.getByTestId("submit-btn");
    expect(submitButton).toBeEnabled();
  });

  it("should call handleSubmit when username and password submitted", async () => {
    mock.onPost("http://localhost:8080/api/signin").reply(200, data);
    render(<SignIn />);
    const userNameInput = screen.getByTestId("username-input");
    fireEvent.change(userNameInput, { target: { value: "tushar" } });
    const passwordInput = screen.getByTestId("password-input");
    fireEvent.change(passwordInput, { target: { value: "tushar" } });
    const submitButton = screen.getByTestId("submit-btn");
    await waitFor(async () => {
      await fireEvent.click(submitButton);
      expect(mockHistoryPush).toHaveBeenCalledTimes(1);
    });
    expect(mockHistoryPush).toHaveBeenCalledWith("/home");
  });

  it("should show inline error messgae when user is unauthorized", async () => {
    mock.onPost("http://localhost:8080/api/signin").reply(401, {});
    render(<SignIn />);
    const userNameInput = screen.getByTestId("username-input");
    fireEvent.change(userNameInput, { target: { value: "tushar" } });
    const passwordInput = screen.getByTestId("password-input");
    fireEvent.change(passwordInput, { target: { value: "tushar" } });
    const submitButton = screen.getByTestId("submit-btn");
    await waitFor(async () => {
      await fireEvent.click(submitButton);
      const linkElement = screen.getByTestId("inline-error");
      expect(linkElement).toBeInTheDocument();
    });
  });
});
