import type { Item } from "../types";

type Props = {
  item: Item;
};

export const ItemCard = ({ item }: Props) => {
  return (
    <div className="bg-white rounded-lg shadow-md p-6 mb-4">
      <h5 className="text-xl font-semibold text-gray-800 mb-3">
        Part Number: {item.partNumber}
      </h5>
      <div className="grid grid-cols-1 md:grid-cols-2 gap-4 text-gray-600">
        <p>Description: {item.description}</p>
        <p>Price: ${item.price?.toFixed(2)}</p>
        <p>Quantity: {item.quantity}</p>
        <p>Purchase Order: {item.purchaseOrder}</p>
        <p>Serial Number: {item.serialNumber}</p>
        <p>Category: {item.category}</p>
        <p>Status: {item.status}</p>
        <p>Repair Order Number: {item.repairOrderNumber}</p>
        <p>Condition: {item.condition}</p>
      </div>
    </div>
  );
};
