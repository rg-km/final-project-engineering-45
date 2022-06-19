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
                        <Link to="/fakultas"><Button buttonStyle="btn--primary" buttonSize="btn--medium">Ketahui Lebih Lanjut</Button></Link>
                    </div>
                </div>
                <div className="main-content-item box-2">
                    <div className="main-content-item--text">
                        <h2 className="heading-content">Persiapan Masuk Perkuliahan</h2>
                        <p className="text-content">Ketahuilah apa saja sih yang perlu kalian persiapkan dengan matang sebelum memasuki dunia perkuliahan ini, pastinya ada banyak hal yang perlu kalian ketahui terlebih dahulu, jadi jangan lewatkan artikel yang kami berikan ini. </p>
                        <Link to="/panduan"><Button buttonStyle="btn--primary" buttonSize="btn--medium">Ketahui Lebih Lanjut</Button></Link>
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