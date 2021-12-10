<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;

class Comentario extends Model
{
    use HasFactory;

    protected $table = 'comentarios';

    protected $fillable =
        [
            'id',
            'id_cliente',
            'descripcion',
            'created_at',
            'updated_at'
        ];
}
