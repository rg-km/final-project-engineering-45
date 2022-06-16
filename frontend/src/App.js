import React from "react";
import { Routes, Route } from "react-router-dom";
import Home from "./content/home/home";
import About from "./content/about/about";
import Fakultas from "./content/fakultas/fakultas";
import FakultasTeknik from "./content/programStudi/fakultasTeknik/fakultasTeknik";

const App = () => {
	return (
		<Routes>
			<Route path="/" element={<Home />} />
			<Route path="/RightWay" element={<About />} />
			<Route path="/fakultas">
				<Route index element={<Fakultas />} />
				<Route path="teknik" element={<FakultasTeknik />} />
			</Route>
		</Routes>
	);
};

export default App;
