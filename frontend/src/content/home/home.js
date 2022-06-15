import React from "react";
import Navbar from "../../component/navbar/navbar";
import HeroSection from "./heroSection/heroSection";
import MainContent from './mainContent/mainContent';
import Footer from "../../component/footer/footer";

const Home = () => {
    return (
        <>
            <Navbar />
            <HeroSection headLine='Welcome to the Right Way' description='Kami adalah teman bagi anda yang akan memberikan berbagai informasi lengkap tentang program studi beserta penjurusan yang ada di dalamnya. Lalu tidak hanya itu, akan ada penjelasan tentang program studi tersebut dan prospek kerja kedepannya seperti apa.' buttonLabel='Ketahui Lebih Lanjut' alt='Hero Section' />
            <MainContent />
            <Footer />
        </>
    )
}

export default Home;