import { createRoot } from "react-dom/client";
import "./index.css";
import { BrowserRouter, Route, Routes } from "react-router";
import Posts from "./posts/Posts";
import "./main.css";
import Navbar from "./components/nav/navbar";

createRoot(document.getElementById("root")!).render(
  <BrowserRouter>
    <Navbar />
    <Routes>
      <Route path="/" element={<Posts />}></Route>
    </Routes>
  </BrowserRouter>,
);
