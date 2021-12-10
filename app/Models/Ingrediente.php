<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;

class Ingrediente extends Model
{
    use HasFactory;

    protected $table = 'ingredientes';

    protected $fillable =
        [
            'id',
            'nombre',
            'descripcion',
            'created_at',
            'updated_at'
        ];
}
