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
			<Route path="RightWay" element={<About />} />
			<Route path="fakultas" element={<Fakultas />}>
				<Route path="teknik">
					<Route index element={<FakultasTeknik />} />
				</Route>
			</Route>
		</Routes>
	);
};

export default App;
