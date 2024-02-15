import { render, screen, waitFor, fireEvent } from "@testing-library/react";
import axios from "axios";
import MockAdapter from "axios-mock-adapter";
import Home from "./Home";
import { ListOfBooks } from "../BooksList/ListOfBooks";

const mock = new MockAdapter(axios);

describe("Search Bar", () => {
  it("Should render Search Bar", () => {
    mock.onGet("http://localhost:8080/api/books").reply(200, {});
    render(<Home />);
    const linkElement = screen.getByTestId("search-bar");
    expect(linkElement).toBeInTheDocument();
  });

  it("Should use data from books api", async () => {
    mock
      .onGet("http://localhost:8080/api/books")
      .reply(200, { books: ListOfBooks });

    await waitFor(async () => {
      await render(<Home />);
      const linkElement = screen.getByTestId("1");
      expect(linkElement).toBeInTheDocument();
    });
  });

  it("Should get data for searched text from books api", async () => {
    mock
      .onGet("http://localhost:8080/api/books?searchText=1")
      .reply(200, { books: [ListOfBooks[0]] });
    render(<Home />);
    const searchInput = screen.getByTestId("search-input");
    fireEvent.change(searchInput, { target: { value: "1" } });
    const searchButton = screen.getByTestId("search-button");
    await fireEvent.click(searchButton);
    await waitFor(async () => {
      const linkElement = screen.getByTestId("1");
      expect(linkElement).toBeInTheDocument();
    });
  });
});
