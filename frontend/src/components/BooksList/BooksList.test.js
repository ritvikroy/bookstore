import { fireEvent, render, screen } from "@testing-library/react";
import { ListOfBooks } from "./ListOfBooks";
import BooksList from "./BooksList";

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

describe("BookList", () => {
  it("should render BooksList", () => {
    render(<BooksList booksList={ListOfBooks} />);
    const linkElement = screen.getByTestId("books-list");
    expect(linkElement).toBeInTheDocument();
  });

  it("should redirect to orderPlaced screen with isOrderPlaced true", () => {
    jest.spyOn(global.Math, "random").mockReturnValue(0.7);
    render(<BooksList booksList={ListOfBooks} />);
    const linkElement = screen.getByTestId("1");
    expect(linkElement).toBeInTheDocument();
    fireEvent.click(linkElement);
    expect(mockHistoryPush).toHaveBeenCalledTimes(1);
    expect(mockHistoryPush).toHaveBeenCalledWith("/orderConfirmation", {
      book: { id: 1, name: "Book 1", price: 100 },
      isOrderPlaced: true,
    });
  });

  it("should redirect to orderPlaced screen with isOrderPlaced false", () => {
    jest.spyOn(global.Math, "random").mockReturnValue(0.3);
    render(<BooksList booksList={ListOfBooks} />);
    const linkElement = screen.getByTestId("1");
    expect(linkElement).toBeInTheDocument();
    fireEvent.click(linkElement);
    expect(mockHistoryPush).toHaveBeenCalledTimes(1);
    expect(mockHistoryPush).toHaveBeenCalledWith("/orderConfirmation", {
      book: { id: 1, name: "Book 1", price: 100 },
      isOrderPlaced: false,
    });
  });
});
