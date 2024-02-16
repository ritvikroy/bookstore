import React, { useState } from "react";
import "./ViewDetails.css";

const ViewDetails = ({ book }) => {
  const [isExpanded, setIdExpanded] = useState(false);

  const handleViewDetails = () => {
    isExpanded ? setIdExpanded(false) : setIdExpanded(true);
  };

  return (
    <>
      <button
        className="view-details-button"
        data-testid={`${book.id}-details`}
        onClick={handleViewDetails}
      >
        {isExpanded ? "Hide Details" : "View Details"}
      </button>
      {isExpanded && (
        <div data-testid={`${book.id}-desc`}>{book.description}</div>
      )}
    </>
  );
};

export default ViewDetails;
