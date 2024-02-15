import { fireEvent, render, screen, waitFor } from "@testing-library/react";
import SearchBar from "./SearchBar";

const mockOnSearch = jest.fn();

describe("Search Bar", () => {
  it("Should render Search Bar", () => {
    render(<SearchBar onSearch={mockOnSearch} />);
    const linkElement = screen.getByTestId("search-bar");
    expect(linkElement).toBeInTheDocument();
  });

  it("should call onSearch on search button click", async () => {
    render(<SearchBar onSearch={mockOnSearch} />);
    const linkElement = screen.getByTestId("search-button");
    fireEvent.click(linkElement);
    const searchInput = screen.getByTestId("search-input");
    await waitFor(async () => {
      await fireEvent.change(searchInput, { target: { value: "the" } });
      expect(mockOnSearch).toHaveBeenCalledTimes(1);
    });
    expect(mockOnSearch).toHaveBeenCalledWith("");
  });
});
