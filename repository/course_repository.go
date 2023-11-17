package repository

import (
	"database/sql"

	"enigmacamp.com/be-lms-university/model/entity"
)

type CourseRepository interface {
	Get(id string) (entity.Course, error)
}

type courseRepository struct {
	db *sql.DB
}

func (c *courseRepository) Get(id string) (entity.Course, error) {
	var course entity.Course
	err := c.db.QueryRow(`SELECT id, course_full_name, course_short_name, description, course_start_date, course_end_date, course_image, created_at, updated_at FROM courses WHERE id = $1`, id).
		Scan(
			&course.Id,
			&course.CourseFullName,
			&course.CourseShortName,
			&course.Description,
			&course.CourseStartDate,
			&course.CourseEndDate,
			&course.CourseImage,
			&course.CreatedAt,
			&course.UpdatedAt,
		)

	if err != nil {
		return entity.Course{}, err
	}

	return course, nil
}

func NewCourseRepository(db *sql.DB) CourseRepository {
	return &courseRepository{db: db}
}
