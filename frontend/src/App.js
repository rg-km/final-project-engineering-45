import React from "react";
import { Routes, Route } from "react-router-dom";
import Home from "./content/home/home";
import About from "./content/about/about";
import Panduan from './content/panduan/Panduan';

const App = () => {
	return (
		<Routes>
			<Route path="/" element={<Home />} />
			<Route path="RightWay" element={<About />} />
			<Route path="panduan">
				<Route index element={<Panduan />} />
			</Route>
		</Routes>
	);
};

export default App;
