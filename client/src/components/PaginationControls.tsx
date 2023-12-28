type Props = {
  currentPage: number;
  totalPages: number;
  onPreviousPage: (currentPage: number) => void;
  onNextPage: (currentPage: number) => void;
  onPageSelect: (selectedPage: number) => void;
};

export const PaginationControls = ({
  currentPage,
  totalPages,
  onPreviousPage,
  onNextPage,
  onPageSelect,
}: Props) => {
  const paginationRange = () => {
    const range = [];
    const delta = 2; // Number of pages before and after the current page
    const left = Math.max(1, currentPage - delta);
    const right = Math.min(totalPages, currentPage + delta);

    for (let i = 1; i <= totalPages; i++) {
      if (i === 1 || i === totalPages || (i >= left && i <= right)) {
        range.push(i);
      } else if (i === left - 1 || i === right + 1) {
        range.push("...");
      }
    }

    return range;
  };

  return (
    <nav aria-label="Page navigation" className="mt-4 flex justify-center">
      <ul className="inline-flex items-center -space-x-px">
        <li>
          <button
            className="py-2 px-3 ml-0 leading-tight text-gray-500 bg-white border border-gray-300 hover:bg-gray-100 hover:text-gray-700"
            onClick={() => onPreviousPage(currentPage)}
            disabled={currentPage <= 1}
          >
            Previous
          </button>
        </li>

        {paginationRange().map((page, index) => (
          <li key={index}>
            {typeof page === "number" ? (
              <button
                className={`py-2 px-3 leading-tight bg-white border border-gray-300 ${
                  currentPage === page
                    ? "text-blue-500 bg-blue-50"
                    : "text-gray-700 hover:bg-gray-100"
                }`}
                onClick={() => onPageSelect(page)}
              >
                {page}
              </button>
            ) : (
              <span className="py-2 px-3 leading-tight bg-white border border-gray-300">
                {page}
              </span>
            )}
          </li>
        ))}

        <li>
          <button
            className="py-2 px-3 leading-tight text-gray-700 bg-white border border-gray-300 hover:bg-gray-100"
            onClick={() => onNextPage(currentPage)}
            disabled={currentPage >= totalPages}
          >
            Next
          </button>
        </li>
      </ul>
    </nav>
  );
};
