export type Item = {
  id: string;
  partNumber: string;
  serialNumber: string;
  purchaseOrder: string;
  description: string;
  category: string;
  price: number;
  quantity: number;
  status: string;
  repairOrderNumber: string;
  condition: string;
  inventoryID: string;
  location: string;
  notes: string;
};

export type ApiUser = {
  id: string;
  name: string;
  email: string;
  role: string;
};

export type FetchArgs = {
  url: string;
  options?: RequestInit;
  includeAuth?: boolean;
};
