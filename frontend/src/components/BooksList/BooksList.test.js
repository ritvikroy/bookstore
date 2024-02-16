import { fireEvent, render, screen, waitFor } from "@testing-library/react";
import { ListOfBooks } from "./ListOfBooks";
import BooksList from "./BooksList";
import MockAdapter from "axios-mock-adapter";
import axios from "axios";
const mock = new MockAdapter(axios);

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

  it("should redirect to orderPlaced screen with isOrderPlaced true", async () => {
    mock.onPost("http://localhost:8080/api/auth/buy-book").reply(200, {});
    render(<BooksList booksList={ListOfBooks} />);
    const linkElement = screen.getByTestId("1");
    await waitFor(async () => {
      await fireEvent.click(linkElement);
      expect(mockHistoryPush).toHaveBeenCalledTimes(1);
    });
    expect(mockHistoryPush).toHaveBeenCalledWith("/orderConfirmation", {
      book: { id: 1, name: "Book 1", price: 100 },
      isOrderPlaced: true,
    });
  });

  it("should redirect to orderPlaced screen with isOrderPlaced false", async () => {
    mock.onPost("http://localhost:8080/api/auth/buy-book").reply(400, {});
    render(<BooksList booksList={ListOfBooks} />);
    const linkElement = screen.getByTestId("1");
    await waitFor(async () => {
      await fireEvent.click(linkElement);
      expect(mockHistoryPush).toHaveBeenCalledTimes(1);
    });
    expect(mockHistoryPush).toHaveBeenCalledWith("/orderConfirmation", {
      book: { id: 1, name: "Book 1", price: 100 },
      isOrderPlaced: false,
    });
  });

  it("should show book details on click of desc", () => {
    render(<BooksList booksList={ListOfBooks} />);
    const viewDetailsButton = screen.getByTestId("1-details");
    fireEvent.click(viewDetailsButton);
    const bookDesc = screen.getByTestId("1-desc");
    expect(bookDesc).toBeInTheDocument();
    fireEvent.click(viewDetailsButton);
    expect(bookDesc).not.toBeInTheDocument();
  });

  it("should not show book details on click of desc 2 times", () => {
    render(<BooksList booksList={ListOfBooks} />);
    const viewDetailsButton = screen.getByTestId("1-details");
    fireEvent.click(viewDetailsButton);
    const bookDesc = screen.getByTestId("1-desc");
    fireEvent.click(viewDetailsButton);
    expect(bookDesc).not.toBeInTheDocument();
  });
});
