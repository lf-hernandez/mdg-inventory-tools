import { Toaster } from "react-hot-toast";

import { Layout } from "./layout/Layout";
import { Home } from "./pages/Home";

function App() {
  return (
    <Layout>
      <Home />
      <Toaster />
    </Layout>
  );
}

export default App;
