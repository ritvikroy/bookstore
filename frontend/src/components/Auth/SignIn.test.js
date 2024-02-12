import { fireEvent, render, screen } from "@testing-library/react";
import SignIn from "./SignIn";

test("should have sign in page", () => {
  render(<SignIn />);
  const linkElement = screen.getByTestId("sign-in");
  expect(linkElement).toBeInTheDocument();
});

test("should render username input",() => {
    render(<SignIn />);
    const linkElement = screen.getByTestId("username-input");
    expect(linkElement).toBeInTheDocument();
});

test("should able to enter username input",() => {
    render(<SignIn />);
    const userNameInput = screen.getByTestId("username-input");
    fireEvent.change(userNameInput, { target: { value: 'tushar' } });
    expect(userNameInput).toHaveValue('tushar');
    expect(userNameInput).toBeInTheDocument();
});


test("should render password input",() => {
    render(<SignIn />);
    const passwordInput = screen.getByTestId("password-input");
    fireEvent.change(passwordInput, { target: { value: 'tushar' } });
    expect(passwordInput).toHaveValue('tushar');
    expect(passwordInput).toBeInTheDocument();
});

test("should render submit btn input",() => {
    render(<SignIn />);
    const linkElement = screen.getByTestId("submit-btn");
    expect(linkElement).toBeInTheDocument();
});