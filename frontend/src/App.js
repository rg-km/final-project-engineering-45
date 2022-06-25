import React from "react";
import { Routes, Route } from "react-router-dom";
import Home from "./content/home/home";
import About from "./content/about/about";
import Fakultas from "./content/fakultas/Fakultas";
import Prodi from "./content/prodi/Prodi";
import ProdiDetail from "./content/prodiDetail/ProdiDetail";
import Panduan from "./content/panduan/Panduan";
import AddProdi from "./content/prodi/add-prodi";

const App = () => {
	return (
		<Routes>
			<Route path="/" element={<Home />} />
			<Route path="RightWay" element={<About />} />
			<Route path="fakultas">
				<Route index element={<Fakultas />} />
				<Route path=":name">
					<Route index element={<Prodi />} />
					<Route path=":prodi" element={<ProdiDetail />} />
				</Route>
			</Route>
			<Route path="panduan">
				<Route index element={<Panduan />} />
			</Route>
			<Route path="/prodi" element={<AddProdi />} />
		</Routes>

	);
};

export default App;
