import React, { useState, useEffect } from "react";
import axios from "axios";
import Navbar from "../../component/navbar/navbar";
import Footer from "../../component/footer/footer";
import { Link, useParams } from "react-router-dom";
import HeaderImage from "../../component/image/background-hero.jpg";
import Image from "../../component/image/prodi-detail.png";
import "./ProdiDetail.css";

const ProdiDetail = () => {
	const [data, setData] = useState([]);
	const [loading, setLoading] = useState(false);
	const [error, setError] = useState(false);
	const { prodi } = useParams();

	useEffect(() => {
		async function getData() {
			setLoading(true);
			setError(false);
			try {
				const url = "http://localhost:8080/api/prodiName?prodi_name=" + prodi;
				const result = await axios.get(url);
				setData(result.data.program_studi);
			} catch (err) {
				setError(true);
			}
			setLoading(false);
		}
		getData();
	}, []);

	return (
		<div>
			<Navbar />
			{loading ? (
				<h1 className="loading">
					<span class="loader"></span>
				</h1>
			) : (
				<div>
					{data.map((item, index) => (
						<div key={index} className="main-container">
							<div
								className="hero-section-prodiDetail"
								style={{
									backgroundImage: `url(${HeaderImage})`,
								}}
							>
								<h1 className="head-hero">{item.prodi_name}</h1>
								<p className="text-hero">
									Penjelasan lengkap mengenai program studi
									yang anda pilih
								</p>
							</div>
							<div className="deskripsi">
								<h1 className="header-deskripsi">
									Apa Itu {item.prodi_name}
								</h1>
								<p className="text-deskripsi">
									{item.deskripsi}
								</p>
							</div>
							<div className="karakter">
								<div className="text-section">
									<h1 className="header-karakter">
										Karakter Yang Sesuai
									</h1>
									<p className="text-karakter">
										{item.karakter}
									</p>
								</div>
								<div className="image-section">
									<img
										src={Image}
										alt="image"
										className="image-karakter"
									/>
								</div>
							</div>
							<div className="box-konten">
								<h1 className="header-matakuliah">
									Matakuliah
								</h1>
								<p className="text-matakuliah">
									{item.mata_kuliah}
								</p>
							</div>
							<div className="box-konten">
								<h1 className="header-prospek">
									Prospek Kerja
								</h1>
								<p className="text-prospek">{item.prospek}</p>
							</div>
						</div>
					))}
				</div>
			)}
			<Footer />
		</div>
	);
};

export default ProdiDetail;
