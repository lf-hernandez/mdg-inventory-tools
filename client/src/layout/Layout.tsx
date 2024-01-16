import { Outlet } from "react-router-dom";
import { Footer } from "./Footer";
import { Sidebar } from "./Sidebar";

export const Layout = () => {
  return (
    <div className="flex flex-col md:flex-row min-h-screen">
      <Sidebar />
      <div className="flex flex-col flex-grow">
        <main className="flex-grow">
          <div className="p-4">
            <Outlet />
          </div>
        </main>
        <Footer />
      </div>
    </div>
  );
};
