package repositorios

import (
	"api/src/modelos"

	"gorm.io/gorm"
)

type Publicacoes struct {
	db *gorm.DB
}

func NovoRepositorioDePublicacoes(db *gorm.DB) *Publicacoes {
	return &Publicacoes{db}
}

func (repositorio Publicacoes) Criar(publicacao modelos.Publicacao) (uint64, error) {
	resultado := repositorio.db.
		Omit("AutorNick", "CriadoEm").
		Create(&publicacao)

	return publicacao.ID, resultado.Error
}

func (repositorio Publicacoes) BuscarPorID(publicacaoID uint64) (modelos.Publicacao, error) {
	var publicacao modelos.Publicacao

	resultado := repositorio.db.
		Table("publicacoes p").
		Select("p.*, u.nick as autor_nick").
		Joins("inner join usuarios u on u.id = p.autor_id").
		Where("p.id = ?", publicacaoID).
		First(&publicacao)

	return publicacao, resultado.Error
}

func (repositorio Publicacoes) BuscarPublicacoes(usuarioID uint64) ([]modelos.Publicacao, error) {
	var publicacoes []modelos.Publicacao

	resultado := repositorio.db.
		Table("publicacoes p").
		Select("p.*, u.nick as autor_nick").
		Joins("INNER JOIN usuarios u ON u.id = p.autor_id").
		Joins("LEFT JOIN seguidores s ON s.usuario_id = p.autor_id").
		Where("p.autor_id = ? OR s.seguidor_id = ?", usuarioID, usuarioID).
		Order("p.criada_em DESC").
		Find(&publicacoes)

	if resultado.Error != nil {
		return nil, resultado.Error
	}

	return publicacoes, nil
}

func (repositorio Publicacoes) BuscarPorUsuario(usuarioID uint64) ([]modelos.Publicacao, error) {
	var publicacoes []modelos.Publicacao

	resultado := repositorio.db.
		Table("publicacoes p").
		Select("p.*, u.nick as autor_nick").
		Joins("INNER JOIN usuarios u ON u.id = p.autor_id").
		Where("p.autor_id = ?", usuarioID).
		Order("p.criada_em DESC").
		Find(&publicacoes)

	if resultado.Error != nil {
		return nil, resultado.Error
	}

	return publicacoes, nil
}

func (repositorio Publicacoes) Atualizar(publicacaoID uint64, publicacao modelos.Publicacao) error {
	resultado := repositorio.db.
		Model(&modelos.Publicacao{}).
		Where("id = ?", publicacaoID).
		Updates(publicacao)
	return resultado.Error
}

func (repositorio Publicacoes) Deletar(publicacaoID uint64) error {
	resultado := repositorio.db.
		Where("id = ?", publicacaoID).
		Delete(&modelos.Publicacao{})
	return resultado.Error
}

func (repositorio Publicacoes) Curtir(publicacaoID uint64) error {
	resultado := repositorio.db.
		Model(&modelos.Publicacao{}).
		Where("id = ?", publicacaoID).
		Updates(map[string]interface{}{"curtidas": gorm.Expr("curtidas + 1")})
	return resultado.Error
}

func (repositorio Publicacoes) Descurtir(publicacaoID uint64) error {
	resultado := repositorio.db.
		Model(&modelos.Publicacao{}).
		Where("id = ? AND curtidas > 0", publicacaoID).
		Update("curtidas", gorm.Expr("curtidas - 1"))

	return resultado.Error
}
