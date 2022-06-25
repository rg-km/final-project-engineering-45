import React from "react";
import { useState } from "react";
import Navbar from "../../component/navbar/navbar";
import Footer from '../../component/footer/footer';
import "./add-prodi.css"

const AddProdi = () => {
    // variable untuk menampung inputan dari form 
    const [ProdiName, setProdiName] = useState('');
    const [FakultasName, setFakultasName] = useState('');
    const [Deskripsi, setDeskripsi] = useState('');
    const [Karakter, setKarakter] = useState('');
    const [MataKuliah, setMataKuliah] = useState('');
    const [Prospek, setProspek] = useState('');

    // function handler submit
    // * buat function menjadi async
    const submit = async (e) => {
        // ! stop reload
        e.preventDefault();

        // * kirim data dari form ke backend melalui endpoint yang sudah kita bikin (http://localhost:8080/api/prodi/add)
        const resp = await fetch('http://localhost:8080/api/prodi/add', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({
                ProdiName,
                FakultasName,
                Deskripsi,
                Karakter,
                MataKuliah,
                Prospek,
            })
        })
    }

    return (
        <>
            <Navbar />
                <div className="mx-5">
            <form>
                <div className="mb-3">
                    <label for="ProdiName" className="form-label">Prodi Name</label>
                    <input type="text" className="form-control" id="ProdiName" placeholder="ProdiName" onChange={e => setProdiName(e.target.value)} required/>
                </div>
                <div className="mb-3">
                    <label for="FakultasName" className="form-label">Fakultas Name</label>
                    <input type="text" className="form-control" id="FakultasName" placeholder="FakultasName" onChange={e => setFakultasName(e.target.value)} required/>
                </div>
                <div className="mb-3">
                    <label for="PDeskripsi" className="form-label">Deskripsi</label>
                    <input type="text" className="form-control" id="Deskripsi" placeholder="Deskripsi" onChange={e => setDeskripsi(e.target.value)} required/>
                </div>
                <div className="mb-3">
                    <label for="Karakter" className="form-label">Karakter</label>
                    <input type="text" className="form-control" id="Karakter" placeholder="Karakter" onChange={e => setKarakter(e.target.value)} required/>
                </div>
                <div className="mb-3">
                    <label for="MataKuliah" className="form-label">Mata Kuliahe</label>
                    <input type="text" className="form-control" id="MataKuliah" placeholder="MataKuliah" onChange={e => setMataKuliah(e.target.value)} required/>
                </div>
                <div className="mb-3">
                    <label for="Prospek" className="form-label">Prospek</label>
                    <input type="text" className="form-control" id="Prospek" placeholder="Prospek" onChange={e => setProspek(e.target.value)} required/>
                </div>
                <button type="submit" className="btn btn-primary d-block mx-auto mb-5">Submit</button>
        
            </form>
            </div>
            <Footer />
        </>
    )
}

export default AddProdi