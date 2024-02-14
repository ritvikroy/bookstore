import { fireEvent, render, screen } from "@testing-library/react";
import SignIn from "./SignIn";

const mockHistoryPush = jest.fn();
jest.mock("react-router", () => ({
  ...jest.requireActual("react-router"),
  useHistory: () => ({
    push: mockHistoryPush,
  }),
}));

test("should have sign in page", () => {
  render(<SignIn />);
  const linkElement = screen.getByTestId("sign-in");
  expect(linkElement).toBeInTheDocument();
});

test("should render username input", () => {
  render(<SignIn />);
  const linkElement = screen.getByTestId("username-input");
  expect(linkElement).toBeInTheDocument();
});

test("should able to enter username input", () => {
  render(<SignIn />);
  const userNameInput = screen.getByTestId("username-input");
  fireEvent.change(userNameInput, { target: { value: "tushar" } });
  expect(userNameInput).toHaveValue("tushar");
});

test("should render password input", () => {
  render(<SignIn />);
  const passwordInput = screen.getByTestId("password-input");
  fireEvent.change(passwordInput, { target: { value: "tushar" } });
  expect(passwordInput).toHaveValue("tushar");
  expect(passwordInput).toBeInTheDocument();
});

// test("should able to enter password input", () => {
//   render(<SignIn />);
//   const passwordInput = screen.getByTestId("password-input");
//   fireEvent.change(passwordInput, { target: { value: "tushar" } });
//   expect(passwordInput).toHaveValue("tushar");
// });

// test("should render submit btn input", () => {
//   render(<SignIn />);
//   const linkElement = screen.getByTestId("submit-btn");
//   expect(linkElement).toBeInTheDocument();
// });

// test("should disable the submit button when username and password is not entered", () => {
//   render(<SignIn />);
//   const submitButton = screen.getByTestId("submit-btn");
//   expect(submitButton).toBeDisabled();
// });

// test("should enabled the submit button when username and password is entered", () => {
//   render(<SignIn />);
//   const userNameInput = screen.getByTestId("username-input");
//   fireEvent.change(userNameInput, { target: { value: "tushar" } });
//   const passwordInput = screen.getByTestId("password-input");
//   fireEvent.change(passwordInput, { target: { value: "tushar" } });
//   const submitButton = screen.getByTestId("submit-btn");
//   expect(submitButton).toBeEnabled();
// });

// test("should call handleSubmit when username and password submitted", () => {
//   render(<SignIn />);
//   const userNameInput = screen.getByTestId("username-input");
//   fireEvent.change(userNameInput, { target: { value: "tushar" } });
//   const passwordInput = screen.getByTestId("password-input");
//   fireEvent.change(passwordInput, { target: { value: "tushar" } });
//   const submitButton = screen.getByTestId("submit-btn");
//   fireEvent.click(submitButton);
//   expect(userNameInput).toHaveValue("");
//   expect(passwordInput).toHaveValue("");
// });
