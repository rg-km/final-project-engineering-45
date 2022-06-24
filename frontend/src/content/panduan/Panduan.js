import React from "react";
import { Link } from "react-router-dom";
import Navbar from "../../component/navbar/navbar";
import Footer from "../../component/footer/footer";
import { MenuItem } from "./MenuItem";
import panduanImage from "../../component/image/panduan.jpg";
import "./Panduan.css";

const Panduan = () => {
	return (
		<>
			<Navbar />
			<div className="box-main">
				<h1>Panduan Masuk Perguruan Tinggi</h1>
				<p>Mari simak panduan yang sudah kita rangkum untuk kalian</p>
				{MenuItem.map((item, index) => (
					<Link key={index} to={item.to} className="link-to">
						<div className="box-panduan">
							<img
								src={panduanImage}
								alt="image"
								className="panduanImage"
							/>
							<div className="text-panduan">
								<h1>{item.title}</h1>
								<p>{item.desc}</p>
							</div>
						</div>
					</Link>
				))}
			</div>
			<Footer />
		</>
	);
};

export default Panduan;
