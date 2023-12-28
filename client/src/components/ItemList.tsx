import { useEffect, useState } from "react";
import { ItemService } from "../services/ItemService";
import type { Item } from "../types";
import { ItemCard } from "./ItemCard";
import { PaginationControls } from "./PaginationControls";

export const ItemList = () => {
  const [items, setItems] = useState<Item[]>([]);
  const [currentPage, setCurrentPage] = useState(1);
  const [totalItems, setTotalItems] = useState(0);
  const totalPages = Math.ceil(totalItems / 10);

  const handlePageSelect = (pageNumber: number) => {
    setCurrentPage(pageNumber);
  };

  const handlePreviousPage = () => {
    if (currentPage > 1) {
      setCurrentPage(currentPage - 1);
    }
  };

  const handleNextPage = () => {
    setCurrentPage(currentPage + 1);
    const sectionTitle = document.querySelector("#itemsListSection");
    if (sectionTitle) {
      sectionTitle.scrollIntoView({ behavior: "smooth" });
    }
  };

  const handleItemUpdate = (updatedItem: Item) => {
    setItems((currentItems) =>
      currentItems.map((item) =>
        item.id === updatedItem.id ? updatedItem : item,
      ),
    );
  };

  useEffect(() => {
    const fetchItems = async () => {
      try {
        const response = await ItemService.getItems(currentPage);
        if (response && response.items) {
          setItems(response.items);
          setTotalItems(response.totalCount);
        } else {
          console.error("Invalid response:", response);
        }
      } catch (error) {
        console.error(error);
      }
    };

    fetchItems();
  }, [currentPage]);

  return (
    <section id="itemsListSection">
      <h2 className="text-2xl font-bold my-4">Items list</h2>
      <div id="itemsList">
        {items.map((item) => (
          <ItemCard key={item.id} item={item} onUpdate={handleItemUpdate} />
        ))}
      </div>
      <PaginationControls
        currentPage={currentPage}
        totalPages={totalPages}
        onPreviousPage={handlePreviousPage}
        onNextPage={handleNextPage}
        onPageSelect={handlePageSelect}
      />
    </section>
  );
};
