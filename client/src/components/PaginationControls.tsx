type Props = {
  currentPage: number;
  onPreviousPage: (currentPage: number) => void;
  onNextPage: (currentPage: number) => void;
};

export const PaginationControls = ({
  currentPage,
  onPreviousPage,
  onNextPage,
}: Props) => {
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

        <li>
          <p
            className={
              "py-2 px-3 leading-tight bg-white border border-gray-300 text-blue-500 bg-blue-50"
            }
          >
            {currentPage}
          </p>
        </li>

        <li>
          <button
            className="py-2 px-3 leading-tight text-gray-700 bg-white border border-gray-300 hover:bg-gray-100"
            onClick={() => onNextPage(currentPage)}
          >
            Next
          </button>
        </li>
      </ul>
    </nav>
  );
};
