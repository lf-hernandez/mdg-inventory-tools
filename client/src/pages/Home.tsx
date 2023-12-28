import { useState } from "react";
import { AddItemForm } from "../components/AddItemForm";
import { ItemList } from "../components/ItemList";
import { SearchForm } from "../components/SearchForm";

const Home = () => {
  const [view, setView] = useState("search");

  return (
    <div className="mx-auto max-w-7xl p-4">
      <div className="flex justify-center space-x-4 mb-6">
        <button
          className={`px-4 py-2 rounded-lg font-semibold transition-colors ${
            view === "search"
              ? "bg-blue-500 text-white"
              : "bg-gray-100 hover:bg-blue-500 hover:text-white"
          }`}
          onClick={() => setView("search")}
        >
          Search Inventory
        </button>
        <button
          className={`px-4 py-2 rounded-lg font-semibold transition-colors ${
            view === "add"
              ? "bg-blue-500 text-white"
              : "bg-gray-100 hover:bg-blue-500 hover:text-white"
          }`}
          onClick={() => setView("add")}
        >
          Add Inventory
        </button>
        <button
          className={`px-4 py-2 rounded-lg font-semibold transition-colors ${
            view === "list"
              ? "bg-blue-500 text-white"
              : "bg-gray-100 hover:bg-blue-500 hover:text-white"
          }`}
          onClick={() => setView("list")}
        >
          View Inventory
        </button>
      </div>
      {view === "search" && <SearchForm />}
      {view === "add" && <AddItemForm />}
      {view === "list" && <ItemList />}
    </div>
  );
};

export default Home;
