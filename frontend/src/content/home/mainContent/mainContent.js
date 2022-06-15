import React from "react";
import penjurusanImage from '../../../component/image/fakultasimage.png';
import materiImage from '../../../component/image/logbookimage.png';
import Button from "../../../component/button/button";
import { Link } from "react-router-dom";
import './mainContent.css';

const MainContent = () => {
    return (
        <div className="main-content">
            <h1 className="primary-heading">Mari, cari tau konten apa yang disediakan di RightWay!</h1>
            <div className="main-content-box">
                <div className="main-content-item">
                    <div className="main-content-item--image">
                        <img src={penjurusanImage} alt="Penjurusan Konten" className="main-image" />
                    </div>
                    <div className="main-content-item--text">
                        <h2 className="heading-content">Fakultas dan Program Studi Sesuai Passion Kalian</h2>
                        <p className="text-content">Panduan untuk kalian yang ingin mengetahui Program Studi apa yang cocok dengan passion kalian, untuk menentukan masa depan terbaik untuk kalian, mulai lah dari memilih hal yang cocok dengan kalian.</p>
                        <Link to="/penjurusan"><Button buttonStyle="btn--primary" buttonSize="btn--medium">Ketahui Lebih Lanjut</Button></Link>
                    </div>
                </div>
                <div className="main-content-item box-2">
                    <div className="main-content-item--text">
                        <h2 className="heading-content">Logbook Kampus Menambah Wawasan</h2>
                        <p className="text-content">Ketahuilah informasi kampus yang ingin kalian jadikan sebagai tempat untuk mengasah skill kalian, demi mencapai masa depan yang kalian inginkan, dengan informasi umum yang diberikan oleh kampus terkait.</p>
                        <Link to="/materi"><Button buttonStyle="btn--primary" buttonSize="btn--medium">Ketahui Lebih Lanjut</Button></Link>
                    </div>
                    <div className="main-content-item--image">
                        <img src={materiImage} alt="Materi Konten" className="main-image" />
                    </div>
                </div>
            </div>
        </div>
    )
}

export default MainContent;