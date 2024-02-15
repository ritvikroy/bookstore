import React from "react";
import { useHistory } from "react-router";
import "./BooksList.css";

const BooksList = ({ booksList }) => {
  const history = useHistory();
  console.log(">>>", booksList);

  const handleBuy = (book) => {
    history.push("/orderConfirmation", {
      book,
      isOrderPlaced: Math.random() > 0.5 ? true : false,
    });
  };

  return (
    <div className="book-list-container" data-testid="books-list">
      <table className="book-table">
        <thead>
          <tr>
            <th>Book Name</th>
            <th>Price</th>
            <th>Action</th>
          </tr>
        </thead>
        <tbody>
          {booksList.map((book) => (
            <tr key={book.id}>
              <td>{book.name}</td>
              <td>&#8377;{book.price}</td>
              <td>
                <button
                  data-testid={`${book.id}`}
                  type="button"
                  onClick={() => handleBuy(book)}
                >
                  Buy
                </button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default BooksList;
