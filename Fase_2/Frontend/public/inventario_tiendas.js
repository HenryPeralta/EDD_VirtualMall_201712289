var elemento = ""
var contador = 0

fetch('http://localhost:3000/cargaproduct')
.then(response => response.json())
.then(data => {
    //elemento.innerHTML = `<p>${data.Datos[0].Departamentos[0].Tiendas[0].Nombre}</p>`;
    for(var i =0; i < (data.Datos).length; i++){
        var departamentos = data.Datos[i].Departamentos
        //console.log(departamentos)
        for(var j=0; j < departamentos.length; j++){
            tiendas = departamentos[j].Tiendas
            //console.log(tiendas)
            for(var k=0; k < tiendas.length; k++){
                var info = tiendas[k].Nombre
                var image = tiendas[k].Logo
                var descrip = tiendas[k].Descripcion
                var contacto = tiendas[k].Contacto
                //var calificacion = tiendas[k].Calificacion
                //contador += 1
                //console.log(info)
                elemento += `<div id='${contador += 1}' class='div_tienda'>
                                <center><h1>Tienda</h1><p>Nombre: ${info} <br> Descripcion: ${descrip} <br> Contacto: ${contacto}</p>                               
                                <img src='${image}' class='image'/> <br>  
                                </center>
                                <br>
                                </div>
                                <br><br>
                                <STYLE type="text/css"> 
                                .div_tienda{width: 500px; margin:0px auto; border: #333 2px solid;} 
                                .image{width: 300px; height: 250px;}
                                </STYLE>`
            }
        }
    }
    //console.log(data)
    document.getElementById("tiendas").innerHTML = elemento;
})
