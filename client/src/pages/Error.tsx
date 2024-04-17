import { ErrorResponse, useRouteError } from "react-router-dom";

export default function Error() {
  // TODO: Find a better way to handle this
  const error = useRouteError() as ErrorResponse & Error;
  console.error(error);

  return (
    <div
      id="error-page"
      className="min-h-screen flex items-center justify-center bg-gray-100"
    >
      <div className="max-w-md w-full p-8 bg-white rounded-lg shadow-lg">
        <h1 className="text-3xl font-bold text-red-600 mb-4">Oops!</h1>
        <p className="text-gray-800 mb-2">
          Sorry, an unexpected error has occurred.
        </p>
        <p className="text-gray-600 mb-4">
          <i>{error.statusText || error.message}</i>
        </p>
        <button
          onClick={() => window.history.back()}
          className="bg-red-600 text-white px-4 py-2 rounded hover:bg-red-700 transition duration-300 ease-in-out"
        >
          Go Back
        </button>
      </div>
    </div>
  );
}
