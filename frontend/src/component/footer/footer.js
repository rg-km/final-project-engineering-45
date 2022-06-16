import React from "react";
import { Link } from "react-router-dom";
import "./footer.css";

const Footer = () => {
    return (
        <footer className="footer">
            <div className="container">
                <div className="row">
                    <h1 className="header-row">Hubungi Kami</h1>
                    <p className='text-row'>telp. 082123231231</p>
                    <p className='text-row'>emai. Rightway@gmail.com</p>
                    <p className='text-row'>alamat. jalan Ruang guru jaya no1</p>
                    <p className='text-row'>kode post. 12345</p>
                </div>
                <div className="row">
                    <h1 className="header-row">Tentang Kami</h1>
                    <p className="text-row">
                        Informasi fakultas dan penjurusan terlengkap dan
                        ter-update, serta logbook dari setiap kampus ternama
                        mengenai informasi umum tentang kampus tersebut, kami
                        siap membantu anda menentukan masa depan untuk anda
                    </p>
                </div>
                <div className="row">
                    <h1 className="header-row">Quote Dari kami</h1>
                    <p className="text-row">
                        "If time passed, and you still do wat you love, then you
                        are on the right way"
                    </p>
                </div>
                <div className="row">
                    <h1 className="header-row">Sosial Media kami</h1>
                    <div className="logo">
                        <Link to="" className="social-media-logo">
                            <i className="fa-brands fa-facebook"></i>Facebook
                        </Link>
                        <Link to="" className="social-media-logo">
                            <i className="fa-brands fa-instagram"></i>Instagram
                        </Link>
                        <Link to="" className="social-media-logo">
                            <i className="fa-brands fa-linkedin"></i>Linkedin
                        </Link>
                        <Link to="" className="social-media-logo">
                            <i className="fa-brands fa-twitter"></i>Twitter
                        </Link>
                    </div>
                </div>
            </div>
            <div className="container">
                <p className="copyright">&#169; 2022, RightWay Website</p>
            </div>
        </footer>
    );
};

export default Footer;
