import React from 'react';
import {Routes, Route} from 'react-router-dom';
import Home from './content/home/home';
import About from './content/about/about';
import Fakultas from './content/fakultas/fakultas';

const App = () => {
	return(
		<Routes>
			<Route path='/' element={<Home />}/>
			<Route path='RightWay' element={<About />} />
			<Route path='fakultas' element={<Fakultas />} />
		</Routes>
	) 
}

export default App;