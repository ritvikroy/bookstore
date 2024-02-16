import { render, screen } from "@testing-library/react";

import App from "./App";

describe("App", () => {
  it("renders home page", () => {
    render(<App />);
    const linkElement = screen.getByTestId("home");
    expect(linkElement).toBeInTheDocument();
  });
});
