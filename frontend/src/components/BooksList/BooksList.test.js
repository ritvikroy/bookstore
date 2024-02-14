import { render, screen } from "@testing-library/react";
import BooksList from "./BooksList";

describe("BookList", () => {
  it("should render BooksList", () => {
    render(<BooksList />);
    const linkElement = screen.getByTestId("books-list");
    expect(linkElement).toBeInTheDocument();
  });
});
