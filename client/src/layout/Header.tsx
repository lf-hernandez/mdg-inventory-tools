import mdgLogo from "../assets/logo.webp";

export const Header = () => {
  return (
    <header className="bg-gray-100 py-4 flex justify-center items-center">
      <img src={mdgLogo} alt="MDG Logo" className="h-20 mr-4" />
      <div className="flex flex-col justify-center h-20">
        <h1 className="text-4xl lg:text-5xl">Inventory Manager</h1>
      </div>
    </header>
  );
};
