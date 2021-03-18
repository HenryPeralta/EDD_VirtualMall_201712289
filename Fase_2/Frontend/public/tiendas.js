var elemento = ""
var contador = 0
var contadorBoton = 0
var departamentos 
var tiendas
var List_tiendas = [];

fetch('http://localhost:3000/Tiendas')
.then(response => response.json())
.then(data => {
    //elemento.innerHTML = `<p>${data.Datos[0].Departamentos[0].Tiendas[0].Nombre}</p>`;
    for(var i =0; i < (data.Datos).length; i++){
        departamentos = data.Datos[i].Departamentos
        //console.log(departamentos)
        for(var j=0; j < departamentos.length; j++){
            tiendas = departamentos[j].Tiendas
            //console.log(tiendas)
            for(var k=0; k < tiendas.length; k++){
                var info = tiendas[k].Nombre
                var image = tiendas[k].Logo
                var descrip = tiendas[k].Descripcion
                var contacto = tiendas[k].Contacto
                var calificacion = tiendas[k].Calificacion
                //contador += 1
                //console.log(info)
                elemento += `<div id='${contador += 1}' class='div_tienda'>
                                <center><h1>Tienda ${contador}</h1><p>Nombre: ${info} <br> Descripcion: ${descrip} <br> Contacto: ${contacto}</p> 
                                <img src='${image}' class='image'/> <br>                                
                                </center>
                                <br>
                                </div>
                                <br><br>
                                <STYLE type="text/css"> 
                                .div_tienda{width: 500px; margin:0px auto; border: #333 2px solid;} 
                                .image{width: 400px; height: 350px;}
                                </STYLE>`
                                console.log(info);
                var local  = new tiendasx(contador, data.Datos[i].Indice, departamentos[j].Nombre, info, descrip, contacto, calificacion, image) 
                List_tiendas.push(local);
                mostrarlista();
            }
        }
    }
    //console.log(data)
    document.getElementById("tienda").innerHTML = elemento;
})

function tiendasx(id, indice, departamento, nombre, descripcion, contacto, calificacion, logo){
    this.id = id;
    this.indice = indice;
    this.departamento = departamento;
    this.nombre = nombre;
    this.descripcion = descripcion;
    this.contacto = contacto;
    this.calificacion = calificacion;
    this.logo = logo;
}

function mostrarlista(){
    var lista ='';
    for(var i=0; i < List_tiendas.length; i++){
        lista += 'id: ' + List_tiendas[i].id +
        ' indice: ' + List_tiendas[i].indice +
        ' departamento: ' + List_tiendas[i].departamento + 
        ' nombre: ' + List_tiendas[i].nombre +
        ' descripcion: ' + List_tiendas[i].descripcion +
        ' contacto: ' + List_tiendas[i].contacto +
        ' calificacion: ' + List_tiendas[i].calificacion + 
        ' logp: ' + List_tiendas[i].logo + '\n';
    }
    console.log(lista);
}

function cargar(){
    var id = parseInt(document.getElementById('tiendaid').value);
    if(!(isNaN(id))){
        for(var i=0; i < List_tiendas.length; i++){
            if(List_tiendas[i].id == id){
                console.log(List_tiendas[i].nombre);
                fetch('http://localhost:3000/cargaproduct', {
                    method: 'POST',
                    headers:{
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify({"Datos":[{"Indice": List_tiendas[i].indice,"Departamentos":[{"Nombre": List_tiendas[i].departamento, "Tiendas":[{"Nombre": List_tiendas[i].nombre, "Descripcion": List_tiendas[i].descripcion, "Contacto": List_tiendas[i].contacto, "Calificacion": List_tiendas[i].calificacion, "Logo": List_tiendas[i].logo}]}]}]})
                })
                //void(window.open('inventario_tiendas.html'));
                location.href="inventario_tiendas.html";
            }
        }
    }
}