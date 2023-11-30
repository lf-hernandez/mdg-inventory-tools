import { AddItemForm } from "../components/AddItemForm";
import { ItemList } from "../components/ItemList";
import { SearchForm } from "../components/SearchForm";

export const Home = () => {
  return (
    <div className="mx-auto max-w-7xl p-4">
      <SearchForm />
      <br />
      <AddItemForm />
      <br />
      <ItemList />
    </div>
  );
};
