import React, { useState } from "react";
import "./SearchBar.css";

const SearchBar = ({ onSearch }) => {
  const [searchText, setSearchText] = useState("");

  const handleChange = (event) => {
    setSearchText(event.target.value);
  };

  const handleSearch = () => {
    onSearch(searchText);
  };

  return (
    <div data-testid="search-bar" className="container">
      <input
        data-testid="search-input"
        className="search-bar-input"
        type="text"
        value={searchText}
        onChange={handleChange}
        placeholder="Search book..."
      />
      <button
        data-testid="search-button"
        className="search-button"
        type="button"
        onClick={handleSearch}
      >
        Search
      </button>
    </div>
  );
};

export default SearchBar;
