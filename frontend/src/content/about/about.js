import React from "react";
import { Link } from "react-router-dom";
import background from "../../component/image/hero-image.jpg";
import Navbar from "../../component/navbar/navbar";
import Footer from '../../component/footer/footer';
import aboutImage from "../../component/image/about1.png";
import Button from "../../component/button/button";
import "./about.css";

const About = () => {
	return (
		<>
			<div
				className="hero-section-about"
				style={{ backgroundImage: `url(${background})` }}
			>
				<div className="hero-container">
					<h1 className="header-about">Welcome To RightWay</h1>
					<p className="text-about">
						Kami akan membantu kalian untuk menentukan masa depan
						yang cocok untuk kalian
					</p>
				</div>
			</div>
			<Navbar />
			<div className="container-about">
				<div className="about-us text">
					<h1>Ingin Tahu Passion Kalian?</h1>
					<p>
						Maka dengan itu kami menyiapkan RightWay yang dimana
						akan membantu kalian untuk menentukan penjurusan yang
						sesuai dengan passion kalian, berbagai pilihan kami
						siapkan, mulai dari fakultas, program studi, dan
						beberapa materi tambahan untuk persiapan kalian memasuki
						dunia perkuliahan
					</p>
					<Link to="/">
						<Button children="Mari Mulai" />
					</Link>
				</div>
				<div className="about-us image">
					<img src={aboutImage} alt="about" />
				</div>
			</div>
			<div className="content-about">
				<h1>Kenapa Pilih RightWay</h1>
				<div className="content-about-section">
					<div className="content-section">
						<h1 className='content header-logo'><i className="fa-solid fa-arrows-spin"></i></h1>
						<h2 className='content'>Mudah & Terintegrasi</h2>
						<p className='content'>
							Sistem infromasi yang di berikan terintegrasi dan
							mudah digunakan
						</p>
					</div>
					<div className="content-section">
						<h1 className='content header-logo'><i className="fa-solid fa-book"></i></h1>
						<h2 className='content'>Informasi Lengkap</h2>
						<p className='content'>
							Fakultas dan penjurusan yang ingin anda cari
							tersedia disini
						</p>
					</div>
					<div className="content-section">
						<h1 className='content header-logo'><i className="fa-solid fa-calendar"></i></h1>
						<h2 className='content'>Mudah Diakses</h2>
						<p className='content'>
							Kami tersedia dalam bentuk website yang mudah di
							akses dan ringan
						</p>
					</div>
				</div>
			</div>
			<div className='content-available'>
				<h1>Konten Yang Kami Berikan</h1>
				<div className='content-tersedia'>
					<div className='box-content'>
						<Link to='/fakultas'>
							<h1 className='logo-content'><i className='fa-solid fa-graduation-cap'></i></h1>
						</Link>
						<p className='text-konten'>Fakultas</p>
					</div>
					<div className='box-content'>
						<Link to='/fakultas'>
							<h1 className='logo-content'><i className='fa-solid fa-building-columns'></i></h1>
						</Link>
						<p className='text-konten'>Program Studi</p>
					</div>
					<div className='box-content'>
						<Link to='/panduan'>
							<h1 className='logo-content'><i className='fa-solid fa-book'></i></h1>
						</Link>
						<p className='text-konten'>Panduan</p>
					</div>
				</div>
			</div>
			<Footer />
		</>
	);
};

export default About;
