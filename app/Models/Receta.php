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
        'tips',
        'calorias',
        'dificultad',
        'cocina',
        'id_cliente',
        'created_at',
        'updated_at'
    ];

    protected $appends =
        [
            'total_favoritos',
            'total_comentarios'
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

    public function clientes_favoritos() 
    {
        return $this->belongsToMany(Cliente::class, 'recetas_favoritas', 'id_receta', 'id_cliente')
            ->select(['id', 'nombre', 'ape_paterno', 'ape_materno']);
    }

    public function cliente()
    {
        return $this->hasOne(Cliente::class, 'id', 'id_cliente' );
    }

    public function getTotalFavoritosAttribute()
    {
        return $this->clientes_favoritos()->count();
    }

    public function getTotalComentariosAttribute()
    {
        return $this->comentarios()->count();
    }
}
