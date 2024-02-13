import React from 'react';
import {ListOfBooks}  from './ListOfBooks';
import './BooksList.css';

const BooksList = () => {
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
          {ListOfBooks.map(book => (
            <tr key={book.id}>
              <td>{book.name}</td>
              <td>Rs {book.price}</td>
              <td>
                <button>Buy</button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default BooksList;