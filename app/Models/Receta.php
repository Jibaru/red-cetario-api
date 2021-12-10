<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;

class Receta extends Model
{
    use HasFactory;
    protected $table = 'recetas';
    protected $fillable =
    [
        'id',
        'titulo',
        'descripcion',
        'tiempo_prep',
        'tiempo_coccion',
        'url_imagen',
        'cocina',
        'id_cliente',
        'created_at',
        'updated_at'
    ];


    public function comentarios()
    {
        return $this->hasMany(Comentario::class, 'id_receta', 'id');
    }

    public function ingredientes()
    {
        return $this->belongsToMany(Ingrediente::class, 'recetas_ingredientes', 'id_receta', 'id_ingrediente')
        ->withPivot('cantidad');
    }

    public function materiales()
    {
        return $this->belongsToMany(Material::class, 'recetas_materiales', 'id_receta', 'id_material');
    }
    
    public function pasos()
    {
        return $this->hasMany(Paso::class, 'id_receta', 'id');
    }
}
