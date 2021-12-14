<?php

use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Support\Facades\Schema;

class CreateRecetasIngredientesTable extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('recetas_ingredientes', function (Blueprint $table) {
            $table->unsignedBigInteger('id_receta');
            $table->foreign('id_receta')
                ->references('id')
                ->on('recetas')
                ->onDelete('cascade');
            $table->unsignedBigInteger('id_ingrediente');
            $table->foreign('id_ingrediente')
                ->references('id')
                ->on('ingredientes')
                ->onDelete('cascade');
            $table->integer('cantidad');
            $table->string('unidad');
        });
    }

    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::dropIfExists('recetas_ingredientes');
    }
}
