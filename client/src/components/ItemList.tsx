import { useEffect, useState } from "react";
import { ItemService } from "../services/ItemService";
import type { Item } from "../types";
import { ItemCard } from "./ItemCard";
import { PaginationControls } from "./PaginationControls";

export const ItemList = () => {
  const [items, setItems] = useState<Item[]>([]);
  const [currentPage, setCurrentPage] = useState(1);

  const handlePreviousPage = () => {
    if (currentPage > 1) {
      setCurrentPage(currentPage - 1);
    }
  };

  const handleNextPage = () => {
    setCurrentPage(currentPage + 1);
  };

  useEffect(() => {
    ItemService.getItems(currentPage).then(setItems).catch(console.error);
  }, [currentPage]);

  return (
    <section id="itemsListSection">
      <h2 className="text-2xl font-bold my-4">Items List</h2>
      <div id="itemsList">
        {items.map((item) => (
          <ItemCard key={item.id} item={item} />
        ))}
      </div>
      <PaginationControls
        currentPage={currentPage}
        onPreviousPage={handlePreviousPage}
        onNextPage={handleNextPage}
      />
    </section>
  );
};
